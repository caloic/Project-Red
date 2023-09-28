package levelManager

import (
	"fmt"
	"ftl/src/enum"
	"math/rand"
	"strconv"
)

func CreateChapter() Chapter {
	return Chapter{Levels: rand.Intn(4) + 7, ActualLevel: 1}
}

func ChangeChapter(c *Chapter) {
	SetLevels(c, rand.Intn(4)+7)
	SetActualLevel(c, 0)
	ChangeActualChapter(c)
}

func ChangeLevel(c *Chapter, zone enum.Zone) {
	if GetActualLevel(c) == GetLevels(c) {
		ChangeChapter(c)
	}
	fmt.Println(c)
	SetLevelZone(c, zone)
	SetActualLevel(c, GetActualLevel(c)+1)
	fmt.Println(zone)
	if GetActualLevel(c) == GetLevels(c)/2 && GetActualChapter(c) == 1 {
		ChangeDifficulty(c)
	} else if GetActualLevel(c) == GetLevels(c)/2 && GetActualChapter(c) == 3 {
		ChangeDifficulty(c)
	} else if GetActualChapter(c) == GetLevels(c) && GetActualChapter(c) == 4 {
		ChangeDifficulty(c)
	}
	fmt.Println(c)
}

func RandomPath(c *Chapter) {
	choice := rand.Intn(4) + 1
	var path []enum.Zone
	var pathchoice int
	for choice > 0 {

		randomIndex := rand.Intn(int(enum.Bonus))
		randomZone := enum.Zone(randomIndex)
		path = append(path, randomZone)
		fmt.Print("Chemin " + strconv.Itoa(choice) + " : ")
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
			fmt.Println("Safe")
		}
		choice -= 1
	}
	fmt.Println(path)
	fmt.Scan(&pathchoice)
	for pathchoice > len(path) || pathchoice <= 0 {
		fmt.Println("Ce chemain n'existe pas")
		fmt.Scan(&pathchoice)
	}
	ChangeLevel(c, path[pathchoice-1])
}
