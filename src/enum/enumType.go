package enum

type Illager int
type Sword int
type Special int
type Zone int
type Armors int

const (
	Pillager    Illager = iota
	Vindicator  Illager = iota
	Evoker      Illager = iota
	Illusioner  Illager = iota
	Ravager     Illager = iota
	Vex         Illager = iota
	EnderDragon Illager = iota
)

const (
	WoodenSword  Sword = iota
	StoneSword   Sword = iota
	IronSword    Sword = iota
	GoldenSword  Sword = iota
	DiamondSword Sword = iota
)

const (
	StoneAxe   Special = iota
	IronAxe    Special = iota
	Golden     Special = iota
	DiamondAxe Special = iota
	Bow        Special = iota
	CrossBow   Special = iota
	Shield     Special = iota
)

const (
	Shop     Zone = iota
	Safe     Zone = iota
	Fight    Zone = iota
	Resource Zone = iota
	Bonus    Zone = iota
)

const (
	Nothing      Armors = iota
	LeatherArmor Armors = iota
	IronArmor    Armors = iota
	GoldenArmor  Armors = iota
	DiamondArmor Armors = iota
)
