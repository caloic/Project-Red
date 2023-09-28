package GameEngine

import (
	"fmt"
	"ftl/src/entity"
	"ftl/src/levelManager"
)

type Arena struct {
	Round       int
	EnemiesList []entity.Illager
}

func BattleArena(c *levelManager.Chapter) {
	arena := Arena{Round: 1}

	mainPlayer := entity.GetPlayerList()[0]

	EnemiesGenerator(arena, c)

	for len(arena.EnemiesList) > 0 && mainPlayer.Hp > 0 {
		fmt.Println("Début du tour:", arena.Round)
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

func EnemiesGenerator(arena Arena, c *levelManager.Chapter) {
	difficulty := levelManager.GetDifficulty(c)

	switch difficulty {
	case 1:
		entity.IllagerCreator(1, arena)
	case 2:
		entity.IllagerCreator(2, arena)
	case 3:
		entity.IllagerCreator(3, arena)
	case 4:
		entity.IllagerCreator(4, arena)
	}
}
