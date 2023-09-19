package entity

import (
	"ftl/src/enum"
)

type Player struct {
	Name      string
	Lifes     int
	Armors    int
	Inventory int
	Money     int
	Armor	  enum.Armor
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

func GetArmors(p *Player) enum.Armor {
	return p.Armors
}

func SetArmors(p *Player, newArmors enum.Armor) {
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
