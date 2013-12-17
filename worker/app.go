package main

import (
	"fmt"

	"github.com/kollinchu/reife/core/trello"
)

func main() {
	fmt.Println(trello.VERSION)

	t := trello.NewTrello("222", "sss")
	fmt.Println(t)
}
