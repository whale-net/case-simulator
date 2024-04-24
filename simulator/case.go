package simulator

import (
	"math/rand"
)

type WeaponCase struct {
	Name string

	WeaponSkins []WeaponSkin
	RarityMap   map[Rarity][]WeaponSkin
}

func (c WeaponCase) Open() WeaponDrop {
	rarity := Rarity(sampleNormal(CumulativeRarityDistribution))
	possibleSkins := c.RarityMap[rarity]

	skinIndex := rand.Intn(len(possibleSkins))
	skinDrop := possibleSkins[skinIndex]

	rawFloat := rand.Float32()
	scaledFloat := skinDrop.MinFloat + (skinDrop.MaxFloat-skinDrop.MinFloat)*rawFloat

	return WeaponDrop{
		WeaponSkin: skinDrop,
		Float:      scaledFloat,
		IsStatTrak: rand.Float32() < 0.10,
	}
}

func NewCase(wsList []WeaponSkin) WeaponCase {
	// initialze with new but return by value
	var wc WeaponCase
	wc.RarityMap = make(map[Rarity][]WeaponSkin)
	wc.WeaponSkins = wsList
	for i := 0; i < len(wc.WeaponSkins); i++ {
		wc.RarityMap[wc.WeaponSkins[i].Rarity] = append(wc.RarityMap[wc.WeaponSkins[i].Rarity], wc.WeaponSkins[i])
	}
	return wc
}

///////////
// Cases //
///////////

// TODO helper function to create
var kilowattCase WeaponCase = NewCase([]WeaponSkin{
	// knife
	{
		Weapon:   Knife,
		Skin:     MadeUpValue,
		Rarity:   Gold,
		MinFloat: 0.0,
		MaxFloat: 1.0,
	},
	// red
	{
		Weapon:   AK47,
		Skin:     Inheritance,
		Rarity:   Red,
		MinFloat: 0.0,
		MaxFloat: 0.80,
	},
	{
		Weapon:   AWP,
		Skin:     ChromeCannon,
		Rarity:   Red,
		MinFloat: 0.0,
		MaxFloat: 1.0,
	},
	// pink
	{
		Weapon:   M4A1S,
		Skin:     BlackLotus,
		Rarity:   Pink,
		MinFloat: 0.0,
		MaxFloat: 0.70,
	},
	{
		Weapon:   Zeusx27,
		Skin:     Olympus,
		Rarity:   Pink,
		MinFloat: 0.0,
		MaxFloat: 0.67,
	},
	{
		Weapon:   USPS,
		Skin:     Jawbreaker,
		Rarity:   Pink,
		MinFloat: 0.0,
		MaxFloat: 1.0,
	},
	// purple
	{
		Weapon:   Glock18,
		Skin:     Block18,
		Rarity:   Purple,
		MinFloat: 0.0,
		MaxFloat: 0.67,
	},
	{
		Weapon:   M4A4,
		Skin:     EtchLord,
		Rarity:   Purple,
		MinFloat: 0.0,
		MaxFloat: 1.0,
	},
	{
		Weapon:   FiveSeven,
		Skin:     Hybrid,
		Rarity:   Purple,
		MinFloat: 0.0,
		MaxFloat: 1.0,
	},
	{
		Weapon:   MP7,
		Skin:     JustSmile,
		Rarity:   Purple,
		MinFloat: 0.0,
		MaxFloat: 1.00,
	},
	{
		Weapon:   SawedOff,
		Skin:     AnalogInput,
		Rarity:   Purple,
		MinFloat: 0.0,
		MaxFloat: 0.62,
	},
	// blue
	{
		Weapon:   MAC10,
		Skin:     LightBox,
		Rarity:   Blue,
		MinFloat: 0.0,
		MaxFloat: 1.0,
	},
	{
		Weapon:   SSG08,
		Skin:     Dezastre,
		Rarity:   Blue,
		MinFloat: 0.0,
		MaxFloat: 1.0,
	},
	{
		Weapon:   XM1014,
		Skin:     Irezumi,
		Rarity:   Blue,
		MinFloat: 0.0,
		MaxFloat: 0.80,
	},
	{
		Weapon:   UMP45,
		Skin:     Motorized,
		Rarity:   Blue,
		MinFloat: 0.0,
		MaxFloat: 0.80,
	},
	{
		Weapon:   Tec9,
		Skin:     Slag,
		Rarity:   Blue,
		MinFloat: 0.0,
		MaxFloat: 0.90,
	},
	{
		Weapon:   DualBerettas,
		Skin:     Hideout,
		Rarity:   Blue,
		MinFloat: 0.0,
		MaxFloat: 0.7,
	},
	{
		Weapon:   Nova,
		Skin:     DarkSigil,
		Rarity:   Blue,
		MinFloat: 0.0,
		MaxFloat: 0.7,
	},
})
