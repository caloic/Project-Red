package entity

import (
	"ftl/src/enum"
)

type Illager struct {
	Type   enum.Illager
	Lifes  int
	attack int
}

func GetType(i *Illager) enum.Illager {
	return i.Type
}

func SetType(i *Illager, newType enum.Illager) {
	i.Type = newType
}

func GetIllagerLifes(i *Illager) int {
	return i.Lifes
}

func SetIllageerLifes(i *Illager, newIllagerLifes int) {
	i.Lifes = newIllagerLifes
}

func GetIllagerAttack(i *Illager) int {
	return i.attack
}

func SetIllagerAttack(i *Illager, newIllagerAttack int) {
	i.attack = newIllagerAttack
}
