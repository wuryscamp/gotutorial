package main

import (
	"fmt"
)

func main() {

	mapPlayer := make(map[int]Player)
	mapPlayer[1] = Player{ID: 1, Name: "Zlatan"}
	mapPlayer[2] = Player{ID: 2, Name: "Higuain"}
	mapPlayer[3] = Player{ID: 3, Name: "Dani Alves"}

	for _, v := range mapPlayer {
		fmt.Println(v.Name)
	}
}

type Player struct {
	ID   int
	Name string
}
