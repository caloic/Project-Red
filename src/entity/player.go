package entity

import (
	"ftl/src/enum"
)

type Player struct {
	Name   string
	Hp     int
	Armors enum.Armors
}

func GetName(p *Player) string {
	return p.Name
}

func SetName(p *Player, newName string) {
	p.Name = newName
}

func GetHp(p *Player) int {
	return p.Hp
}

func SetHp(p *Player, newHp int) {
	p.Hp = newHp
}

func GetArmors(p *Player) enum.Armors {
	return p.Armors
}

func SetArmors(p *Player, newArmors enum.Armors) {
	p.Armors = newArmors
}
