package main

import (
	"fmt"
	"ftl/src/levelManager"
)

func main() {
	c := levelManager.CreateChapterGen()
	fmt.Println(c)
	levelManager.ChangeChapter(&c)
	fmt.Println(c)
	levelManager.RandomZone(&c)
}
