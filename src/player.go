package main

type Player struct {
	Name string
	Lifes int
	Armors int
	Inventory int
	Money int
}

func (p Player) GetName() string {
	return p.Name
}

func (p Player) SetName(name string) {
	p.Name = name
}