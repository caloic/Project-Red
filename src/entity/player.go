package entity

type Player struct {
	Name      string
	Lifes     int
	Armors    int
	Inventory map[string]int
	Money     int
}

func (p Player) GetName() string {
	return p.Name
}

func (p Player) SetName(name string) {
	p.Name = name
}

func (p Player) GetLifes() int {
	return p.Lifes
}

func (p Player) GetArmors() int {
	return p.Armors
}

// Taille de l'inventaire du joueur
func (p Player) GetInventory() map {
	return p.Inventory
}

func (p Player) GetMoney() int {
	return p.Money
}

func (p Player) SetMoney(money int) {
	p.Money = money
}
