package entity

import (
	"ftl/src/enum"
)

type Player struct {
	Name         string
	Lifes        int
	Armors       enum.Armors
	HoldingSword map[enum.Sword]int
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

func GetArmors(p *Player) enum.Armors {
	return p.Armors
}

func SetArmors(p *Player, newArmors enum.Armors) {
	p.Armors = newArmors
}

func GetHoldingSword(p *Player) map[enum.Sword]int {
	return p.HoldingSword
}

func SetHoldingSword(p *Player, newHoldingSword map[enum.Sword]int) {
	p.HoldingSword = newHoldingSword
}
