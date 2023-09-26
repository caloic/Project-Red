package GameEngine

import "ftl/src/entity"

type Arena struct {
	Turn int
}

func BattleArena() {
	mainPlayer := entity.GetPlayerList()[0]
	for mainPlayer.Hp > 0 {

	}
}
