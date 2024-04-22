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

func (r Rarity) String() string {
	switch r {
	case Blue:
		return fmt.Sprint("Blue")
	case Purple:
		return fmt.Sprint("Purple")
	case Pink:
		return fmt.Sprint("Pink")
	case Red:
		return fmt.Sprint("Red")
	case Gold:
		return fmt.Sprint("Gold")
	default:
		return fmt.Sprint("Unknown")
	}
}

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
