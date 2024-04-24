package simulator

import "fmt"

// type of weapon
type Weapon int

// weapon + skin
type WeaponSkin struct {
	Weapon Weapon
	Skin   Skin
	Rarity Rarity

	MinFloat float32
	MaxFloat float32
}

func (ws WeaponSkin) String() string {
	return fmt.Sprintf("%v %v (%v)", ws.Weapon, ws.Skin, ws.Rarity)
}

// TODO - GetQualities

// Instance of WeaponSkin
type WeaponDrop struct {
	WeaponSkin WeaponSkin
	Float      float32
	IsStatTrak bool
}

func (wd WeaponDrop) GetQuality() Quality {
	// technically wrong order for speed, if it were to even matter
	if wd.Float <= 0.07 {
		return FactoryNew
	} else if wd.Float <= 0.15 {
		return MinimalWear
	} else if wd.Float <= 0.38 {
		return FieldTested
	} else if wd.Float <= 0.45 {
		return WellWorn
	}
	return BattleScarred
}

func (wd WeaponDrop) String() string {
	stattrak := ""
	if wd.IsStatTrak {
		stattrak = "StatTrak"
	}
	return fmt.Sprintf("%v %v - %v (%v)", stattrak, wd.WeaponSkin, wd.GetQuality(), wd.Float)
}

// At bottom of file because this rarely changes
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
