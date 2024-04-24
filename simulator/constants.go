package simulator

import "fmt"

type Rarity int

const (
	Blue Rarity = iota
	Purple
	Pink
	Red
	Gold
)

// until I can figure out a better way to do this, I'm going to do this a lot
// at least editor tools make this effortless
func (r Rarity) String() string {
	switch r {
	case Blue:
		return fmt.Sprint("Blue - Mil-Spec")
	case Purple:
		return fmt.Sprint("Purple - Restricted")
	case Pink:
		return fmt.Sprint("Pink - Classified")
	case Red:
		return fmt.Sprint("Red - Covert")
	case Gold:
		return fmt.Sprint("Gold - Rare Special")
	default:
		return fmt.Sprint("Unknown")
	}
}

var RarityDistribution Distribution = Distribution{0.7992, 0.1598, 0.0320, 0.0064, 0.0026}
var CumulativeRarityDistribution Distribution = RarityDistribution.toCumulative(false)

// TODO - this should be lookup to mapping of floats per weapon
// something more complex
// although this is probably not the right thing for it
type Quality int

const (
	BattleScarred = iota
	WellWorn
	FieldTested
	MinimalWear
	FactoryNew
)

func (q Quality) String() string {
	switch q {
	case BattleScarred:
		return fmt.Sprint("Battle Scarred")
	case WellWorn:
		return fmt.Sprint("Well Worn")
	case FieldTested:
		return fmt.Sprint("Field Tested")
	case MinimalWear:
		return fmt.Sprint("Minimal Wear")
	case FactoryNew:
		return fmt.Sprint("Factory New")
	default:
		return fmt.Sprint("Unknown")
	}
}
