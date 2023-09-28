package entity

import "ftl/src/enum"

var playerList []Player

func CreatePlayer() {
	listLen := len(playerList)
	switch listLen {
	case 0:
		playerList = append(playerList, Player{Hp: 20, Armors: enum.Nothing, Name: "Player"})
	case 1:
		playerList = append(playerList, Player{Hp: 20, Armors: enum.Nothing, Name: "Soldier1"})
	case 2:
		playerList = append(playerList, Player{Hp: 20, Armors: enum.Nothing, Name: "Soldier2"})
	case 3:
		playerList = append(playerList, Player{Hp: 20, Armors: enum.Nothing, Name: "Soldier3"})
	}
}

func GetPlayerList() []Player {
	return playerList
}

func PlayerAttack(e Illager, player Player) {
	if {

	}
}
