package main

import (
	"fmt"
	"ftl/src/mob"
)

func main() {
	user := mob.Player{Name: "JK"}
	fmt.Print(user.GetName())
}
