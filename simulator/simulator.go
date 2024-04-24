package simulator

import (
	"fmt"
	"maps"
	"strings"
	"sync"
)

// TODO more sophisticated results
//
//	type SimulationResult struct {
//		DropResult map[WeaponDrop]int
//	}
type SimulationResult struct {
	weaponSkinTotals map[WeaponSkin]int
}

func NewSimulationResult() *SimulationResult {
	return &SimulationResult{
		weaponSkinTotals: make(map[WeaponSkin]int),
	}
}

func (sr *SimulationResult) merge(other *SimulationResult) {
	// left join sum
	for wd := range sr.weaponSkinTotals {
		if _, ok := other.weaponSkinTotals[wd]; ok {
			sr.weaponSkinTotals[wd] += other.weaponSkinTotals[wd]
		}
	}

	// pick up missing right side - doesn't overwrite, so it will carry through any extra on right side
	maps.Copy(sr.weaponSkinTotals, other.weaponSkinTotals)
}

func (sr *SimulationResult) processWeaponDrop(wd *WeaponDrop) {
	sr.weaponSkinTotals[wd.WeaponSkin]++
}

func (sr SimulationResult) String() string {
	var sb strings.Builder
	rarityCount := make(map[Rarity]int)
	totalCount := 0

	sb.WriteString("Total Drops:\n")
	for wd, count := range sr.weaponSkinTotals {
		rarityCount[wd.Rarity] += count
		totalCount += count
		sb.WriteString(fmt.Sprintf("\t[%v] %v times\n", wd, count))
	}

	fmt.Println("Rarity:")
	for rarity, count := range rarityCount {
		rarityChance := float32(count) / float32(totalCount)
		sb.WriteString(fmt.Sprintf("[%v] %v times (%v%%)\n", rarity, count, rarityChance*100))
	}

	return sb.String()
}

// these are global but instantiated multiple times unfortunately
var chances = make(map[Rarity]float32)
var cumulative_chances = make(map[Rarity]float32)

func Simulate(batchCount int, batchSize int) {
	if batchCount <= 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(batchCount)

	// definitely shouldn't use a channel this big, but...
	resultChannel := make(chan *SimulationResult, batchCount)

	for i := 0; i < batchCount; i++ {
		go simulate(batchSize, &wg, resultChannel)
	}

	wg.Wait()
	close(resultChannel)

	// merge all reuslts into the main result
	mainResult := <-resultChannel
	for otherResult := range resultChannel {
		mainResult.merge(otherResult)
	}

	fmt.Println("FINAL RESULT")
	fmt.Println(fmt.Sprintf("Total Cases Opened: %v", batchCount*batchSize))
	fmt.Println(mainResult)
}

func simulate(batchSize int, wg *sync.WaitGroup, resultChannel chan *SimulationResult) {
	defer wg.Done()

	if batchSize <= 0 {
		return
	}

	result := NewSimulationResult()
	for i := 0; i < batchSize; i++ {
		wd := kilowattCase.Open()
		result.processWeaponDrop(&wd)
	}

	resultChannel <- result
	fmt.Println(result)
}
