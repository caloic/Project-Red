package enum

type Illager int
type Sword int
type Axe int
type Zone int
type Armor int

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
	StoneSword   Sword = iota
	IronSword    Sword = iota
	GoldenSword  Sword = iota
	DiamondSword Sword = iota
)

const (
	StoneAxe   Axe = iota
	IronAxe    Axe = iota
	Golden     Axe = iota
	DiamondAxe Axe = iota
)

const (
	Shop     Zone = iota
	Safe     Zone = iota
	Fight    Zone = iota
	Resource Zone = iota
	Bonus    Zone = iota
)

const (
	LeatherArmor Armor = iota
	IronArmor Armor = iota
	GoldenArmor Armor = iota
	DiamondArmor Armor = iota
)
