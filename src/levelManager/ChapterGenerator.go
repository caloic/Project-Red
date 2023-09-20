package levelManager

import (
	"fmt"
	"ftl/src/enum"
	"math/rand"
)

func CreateChapterGen() Chapter {
	return Chapter{Levels: rand.Intn(4) + 7, ActualLevel: 0}
}

func ChangeChapter(c *Chapter) {
	SetLevels(c, rand.Intn(4)+7)
	SetActualLevel(c, 0)
}

func RandomZone(c *Chapter) {
	choice := rand.Intn(4) + 1
	for choice > 0 {

		randomIndex := rand.Intn(int(enum.Bonus))
		randomZone := enum.Zone(randomIndex)
		switch randomZone {
		case enum.Safe:
			fmt.Println("Safe")
		case enum.Shop:
			fmt.Println("Shop")
		case enum.Resource:
			fmt.Println("Safe")
		case enum.Fight:
			fmt.Println("Fight")
		case enum.Bonus:
			fmt.Println("Bonus")
		}
		choice -= 1
	}
}
