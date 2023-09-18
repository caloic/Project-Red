package entity

type Player struct {
	Name      string
	Lifes     int
	Armors    int
	Inventory int
	Money     int
	//Ressource
	Wood    int
	Stone   int
	Coal    int
	Iron    int
	Gold    int
	Diamond int
	emerald int
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
func (p Player) GetInventory() int {
	return p.Inventory
}

func (p Player) GetMoney() int {
	return p.Money
}

func (p Player) SetMoney(money int) {
	p.Money = money
}
