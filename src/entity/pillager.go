package entity

type Pillager struct {
	Type string
	Lifes int 
	Armors int
	Inventory int 
}

func (p Pillager) GetName() string {
	return p.Name
}

func (p Pillager) SetName() string {
	p.Name = name
}

func (p Villager) GetLifes() string {
	return p.Lifes
}

func (p Villager) GetArmors() string {
	return p.Armors
}

func (p Villager) GetInventory() string {
	return p.Inventory
}