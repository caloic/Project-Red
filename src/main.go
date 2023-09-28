package main

import "ftl/src/levelManager"

func main() {
	chap := levelManager.CreateChapter()
	levelManager.RandomPath(&chap)
}
