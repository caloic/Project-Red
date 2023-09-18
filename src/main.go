package main

import (
	"fmt"
	"ftl/src/entity"
)

func main() {
	user := entity.Player{Name: "JK"}
	fmt.Print(user.GetName())
}
