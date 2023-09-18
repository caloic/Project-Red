package main

import (
	"fmt"
	"ftl/src/player"
)

func main() {
	user := player.Player{Name: "JK"}
	fmt.Print(user.GetName())
}
