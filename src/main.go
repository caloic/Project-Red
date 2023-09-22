package main

import (
	"fmt"
	"ftl/src/levelManager"
	"ftl/src/shop"
)

func main() {
	c := levelManager.CreateChapterGen()
	fmt.Println(c)
	levelManager.ChangeChapter(&c)
	fmt.Println(c)
	levelManager.RandomZone(&c)
	shop.Shop()
}
