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

type Weapon int

const (
	Knife Weapon = iota
	AK47
	AWP
	M4A1S
	Zeusx27
	USPS
	Glock18
	M4A4
	FiveSeven
	MP7
	SawedOff
	MAC10
	SSG08
	XM1014
	UMP45
	Tec9
	DualBerettas
	Nova
)

func (w Weapon) String() string {
	switch w {
	case AK47:
		return "AK47"
	case AWP:
		return "AWP"
	case M4A1S:
		return "M4A1S"
	case Zeusx27:
		return "Zeusx27"
	case USPS:
		return "USPS"
	case Glock18:
		return "Glock18"
	case M4A4:
		return "M4A4"
	case FiveSeven:
		return "FiveSeven"
	case MP7:
		return "MP7"
	case SawedOff:
		return "SawedOff"
	case MAC10:
		return "MAC10"
	case SSG08:
		return "SSG08"
	case XM1014:
		return "XM1014"
	case UMP45:
		return "UMP45"
	case Tec9:
		return "Tec9"
	case DualBerettas:
		return "DualBerettas"
	case Nova:
		return "Nova"
	default:
		return "Unknown"
	}
}

type Skin int

const (
	MadeUpValue Skin = iota
	Inheritance
	ChromeCannon
	BlackLotus
	Olympus
	Jawbreaker
	Block18
	EtchLord
	Hybrid
	JustSmile
	AnalogInput
	LightBox
	Dezastre
	Irezumi
	Motorized
	Slag
	Hideout
	DarkSigil
)

func (s Skin) String() string {
	switch s {
	case MadeUpValue:
		return "Developer Value"
	case Inheritance:
		return "Inheritance"
	case ChromeCannon:
		return "ChromeCannon"
	case BlackLotus:
		return "BlackLotus"
	case Olympus:
		return "Olympus"
	case Jawbreaker:
		return "Jawbreaker"
	case Block18:
		return "Block18"
	case EtchLord:
		return "EtchLord"
	case Hybrid:
		return "Hybrid"
	case JustSmile:
		return "JustSmile"
	case AnalogInput:
		return "AnalogInput"
	case LightBox:
		return "LightBox"
	case Dezastre:
		return "Dezastre"
	case Irezumi:
		return "Irezumi"
	case Motorized:
		return "Motorized"
	case Slag:
		return "Slag"
	case Hideout:
		return "Hideout"
	case DarkSigil:
		return "DarkSigil"
	default:
		return "Unknown"
	}
}
