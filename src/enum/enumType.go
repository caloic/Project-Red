package enum

type Illager int
type Sword int
type Axe int

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
