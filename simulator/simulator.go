package simulator

import (
	"fmt"
	"maps"
	"math/rand"
	"sync"
)

type WeaponDrop struct {
	Rarity  Rarity
	Quality Quality
	Name    string
}

func (wd WeaponDrop) String() string {
	return fmt.Sprintf("%v (%v) - %v", wd.Name, wd.Rarity, wd.Quality)
}

//	type SimulationResult struct {
//		DropResult map[WeaponDrop]int
//	}
type SimulationResult map[WeaponDrop]int

func (self *SimulationResult) merge(other *SimulationResult) {
	// left join sum
	for wd, _ := range *self {
		if _, ok := (*other)[wd]; ok {
			(*self)[wd] += (*other)[wd]
		}
	}

	// pick up missing right side - won't overwrite so will carry through any extra
	maps.Copy(*self, *other)
}

// these are global but instantiated multiple times unfortunately
var chances = make(map[Rarity]float32)
var cumulative_chances = make(map[Rarity]float32)

func SimulateMany(count int) {
	if count <= 0 {
		return
	}

	// how to initialize this outside this function
	chances[Blue] = 0.7992
	chances[Purple] = 0.1598
	chances[Pink] = 0.0320
	chances[Red] = 0.0064
	chances[Gold] = 0.0026

	// how to do this idiomatically
	cumulative_chances[Blue] = chances[Blue]
	cumulative_chances[Purple] = cumulative_chances[Blue] + chances[Purple]
	cumulative_chances[Pink] = cumulative_chances[Purple] + chances[Pink]
	cumulative_chances[Red] = cumulative_chances[Pink] + chances[Red]
	cumulative_chances[Gold] = cumulative_chances[Red] + chances[Gold]

	var wg sync.WaitGroup
	wg.Add(count)

	// definitely shouldn't use a channel this big, but...
	resultChannel := make(chan SimulationResult, count)

	for i := 0; i < count; i++ {
		go Simulate(&wg, resultChannel)
	}

	wg.Wait()
	close(resultChannel)

	// merge all reuslts into the main result
	mainResult := <-resultChannel
	for otherResult := range resultChannel {
		mainResult.merge(&otherResult)
	}

	fmt.Println("FINAL RESULT")
	printResults(&mainResult)

	return
}

func Simulate(wg *sync.WaitGroup, resultChannel chan SimulationResult) SimulationResult {
	if wg != nil {
		defer wg.Done()
	} else {
		// assuming if wg is nil that not called from SimulateMany - weak assumption but easy
		// how to initialize this outside this function
		chances[Blue] = 0.7992
		chances[Purple] = 0.1598
		chances[Pink] = 0.0320
		chances[Red] = 0.0064
		chances[Gold] = 0.0026

		// how to do this idiomatically
		cumulative_chances[Blue] = chances[Blue]
		cumulative_chances[Purple] = cumulative_chances[Blue] + chances[Purple]
		cumulative_chances[Pink] = cumulative_chances[Purple] + chances[Pink]
		cumulative_chances[Red] = cumulative_chances[Pink] + chances[Red]
		cumulative_chances[Gold] = cumulative_chances[Red] + chances[Gold]
	}

	var dropResults = make(SimulationResult)

	var num_drops = 100_000_000

	fmt.Println(fmt.Sprintf("Simulating %v drops", num_drops))

	for i := 0; i < num_drops; i++ {
		var drop = simulate()
		//fmt.Println(drop)
		dropResults[drop]++
	}

	if resultChannel != nil {
		resultChannel <- dropResults
	}

	//printResults(dropResults)

	return dropResults
}

func simulate() WeaponDrop {

	var drop WeaponDrop

	if random := rand.Float32(); random <= cumulative_chances[Blue] {
		drop.Rarity = Blue
	} else if random <= cumulative_chances[Purple] {
		drop.Rarity = Purple
	} else if random <= cumulative_chances[Pink] {
		drop.Rarity = Pink
	} else if random <= cumulative_chances[Red] {
		drop.Rarity = Red
	} else if random <= cumulative_chances[Gold] {
		drop.Rarity = Gold
	}

	// TODO pick something
	// CaseDrop interface
	drop.Name = "no name"
	drop.Quality = FactoryNew

	return drop
}

func printResults(results *SimulationResult) {
	var rarityCount = make(map[Rarity]int)
	totalCount := 0

	fmt.Println("Drops:")
	for wd, count := range *results {
		rarityCount[wd.Rarity] += count
		totalCount += count
		fmt.Println(fmt.Sprintf("\t[%v] %v times", wd, count))
	}

	fmt.Println("Rarity:")
	for rarity, count := range rarityCount {
		rarityChance := float32(count) / float32(totalCount)
		fmt.Println(fmt.Sprintf("[%v] %v times (%v%%)", rarity, count, rarityChance*100))
	}
}
