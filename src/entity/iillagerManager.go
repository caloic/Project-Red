package entity

import (
	"ftl/src/GameEngine"
)

func IllagerCreator(i int, f GameEngine.Arena) {
	loop := 1
	for loop < i {
		f.EnemiesList = append(f.EnemiesList, Illager{Type: 1, Lifes: 20, attack: 2})
	}
}
