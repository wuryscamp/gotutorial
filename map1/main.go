package main

import (
	"fmt"
)

func main() {

	var mapPerson map[int]string
	mapPerson = make(map[int]string)

	mapPerson[1] = "Alex"
	mapPerson[2] = "Ben"
	mapPerson[3] = "Den"
	mapPerson[4] = "Wury"

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}

	wury, ok := mapPerson[4]
	if !ok {
		fmt.Println("Wury tidak ada di mapPerson")
	} else {
		fmt.Println(wury)
	}

}
