package entity

import "ftl/src/enum"

var playerList []Player

func CreatePlayer() {
	playerList = append(Player{Lifes: 20, Armors: enum.Nothing, Name: "Player"})
}

func LooseHealth(p Player, pv int) {
	p.Lifes = p.Lifes - pv
}
