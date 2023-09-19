package entity

type Player struct {
	Name      string
	Lifes     int
	Armors    int
	Inventory int
	Money     int
}

func GetName(p *Player) string {
	return p.Name
}

func SetName(p *Player, newName string) {
	p.Name = newName
}

func GetLifes(p *Player) int {
	return p.Lifes
}

func SetLifes(p *Player, newLifes int) {
	p.Lifes = newLifes
}

func GetArmors(p *Player) int {
	return p.Armors
}

func SetArmors(p *Player, newArmors int) {
	p.Armors = newArmors
}

func GetInventory(p *Player) int {
	return p.Inventory
}

func GetMoney(p *Player) int {
	return p.Money
}

func SetMoney(p *Player, newMoney int) {
	p.Money = newMoney
}
