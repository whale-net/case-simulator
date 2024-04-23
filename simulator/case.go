package simulator

import (
	"fmt"
	"math/rand"
)

// Instance of WeaponSkin
type WeaponDrop struct {
	WeaponSkin WeaponSkin
	Float      float32
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
	return fmt.Sprintf("%v - %v (%v)", wd.WeaponSkin, wd.GetQuality(), wd.Float)
}

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

type Case struct {
	Name string
	// TODO make list and add helper functions
	SkinPool map[Rarity][]WeaponSkin
}

func (c Case) Open() WeaponDrop {
	rarity := Rarity(sampleNormal(CumulativeRarityDistribution))
	possibleSkins := c.SkinPool[rarity]

	skinIndex := rand.Intn(len(possibleSkins))
	skinDrop := possibleSkins[skinIndex]

	rawFloat := rand.Float32()
	scaledFloat := skinDrop.MinFloat + (skinDrop.MaxFloat-skinDrop.MinFloat)*rawFloat

	return WeaponDrop{
		WeaponSkin: skinDrop,
		Float:      scaledFloat,
	}
}

///////////
// Cases //
///////////

// TODO helper function to create
var kilowattCase Case = Case{
	Name: "Kilowatt",
	SkinPool: map[Rarity][]WeaponSkin{
		// TODO Fill out Knives properly
		Gold: {
			{
				Weapon:   Knife,
				Skin:     MadeUpValue,
				Rarity:   Gold,
				MinFloat: 0.0,
				MaxFloat: 1.0,
			},
		},
		Red: {
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
		},
		Pink: {
			{
				Weapon:   M4A1S,
				Skin:     BlackLotus,
				Rarity:   Pink,
				MinFloat: 0.0,
				MaxFloat: 0.70,
			},
		},
		Purple: {
			{
				Weapon:   Glock18,
				Skin:     Block18,
				Rarity:   Purple,
				MinFloat: 0.0,
				MaxFloat: 0.67,
			},
		},
		Blue: {
			{
				Weapon:   MAC10,
				Skin:     LightBox,
				Rarity:   Blue,
				MinFloat: 0.0,
				MaxFloat: 1.0,
			},
		},
	},
}
