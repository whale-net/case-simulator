package simulator

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
