package inventory

type Inventory struct {
	InventorySpace int
	Wood int
	Stone int
	Iron int
	Gold int
	Diamond int
	Emerald	int
	StoneSword int
	IronSword int
	GoldenSword int
	DiamondSword int
	TotemUndying int
}

func GetInventorySpace(i *Inventory) int {
	return i.InventorySpace
}

func SetInventorySpace(i *Inventory, newInventorySpace int) {
	i.InventorySpace = newInventorySpace
}

func GetWood(i *Inventory) int {
	return i.Wood
}

func SetWood(i *Inventory, newWood int) {
	i.Wood = newWood
}

func GetStone(i *Inventory) int {
	return i.Stone
}

func SetStone(i *Inventory, newStone int) {
	i.Stone = newStone
}

func GetIron(i *Inventory) int {
	return i.Iron
}

func SetIron(i *Inventory, newIron int) {
	i.Iron = newIron
}

func GetGold(i *Inventory) int {
	return i.Gold
}

func SetGold(i *Inventory, newGold int) {
	i.Gold = newGold
}

func GetDiamond(i *Inventory) int {
	return i.Diamond
}

func SetDiamond(i *Inventory, newDiamond int) {
	i.Diamond = newDiamond
}

func GetEmerald(i *Inventory) int {
	return i.Emerald
}

func SetEmerald(i *Inventory, newEmerald int) {
	i.Emerald = newEmerald
}

func GetStoneSword(i *Inventory) int {
	return i.StoneSword
}

func SetStoneSword(i *Inventory, newStoneSword int) {
	i.StoneSword = newStoneSword
}

func GetIronSword(i *Inventory) int {
	return i.IronSword
}

func SetIronSword(i *Inventory, newIronSword int) {
	i.IronSword = newIronSword
}

func GetGoldenSword(i *Inventory) int {
	return i.GoldenSword
}

func SetGoldenSword(i *Inventory, newGoldenSword int) {
	i.GoldenSword = newGoldenSword
}

func GetDiamondSword(i *Inventory) int {
	return i.DiamondSword
}

func SetDiamondSword(i *Inventory, newDiamondSword int) {
	i.DiamondSword = newDiamondSword
}

func GetTotemUndying(i *Inventory) int {
	return i.TotemUndying
}

func SetTotemUndying(i *Inventory, newTotemUndying int) {
	i.TotemUndying = newTotemUndying
}