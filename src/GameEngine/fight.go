package GameEngine

import (
	"fmt"
	"ftl/src/entity"
)

type Arena struct {
	Round       int
	EnemiesList []entity.Illager
}

func BattleArena() {
	arena := Arena{Round: 1}

	mainPlayer := entity.GetPlayerList()[0]

	for len(arena.EnemiesList) > 0 && mainPlayer.Hp > 0 {
		fmt.Println("Round:", arena.Round)
		arena.Round++

		if mainPlayer.Hp <= 0 {
			break
		}
	}

	if len(arena.EnemiesList) == 0 {
		fmt.Println("Tous les ennemis ont été vaincus.")
	} else {
		fmt.Println("Le joueur principal a été vaincu.")
	}
}
