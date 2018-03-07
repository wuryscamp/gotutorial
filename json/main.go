package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	player1 := Player{ID: 1, Name: "Diego Costa"}

	p1, err := json.Marshal(player1)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(p1))

	fmt.Println("===========================")

	data := []byte(`{"id": 1, "name": "Gary Cahlil"}`)

	var p2 Player

	err = json.Unmarshal(data, &p2)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(p2.ID)
	fmt.Println(p2.Name)

}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
