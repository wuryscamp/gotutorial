package main

import (
	"fmt"
)

func main() {

	x := 100

	// expression switch
	switch x {
	case 60:
		fmt.Println("NILAI = C")
	case 80:
		fmt.Println("NILAI = B")
	case 100:
		fmt.Println("NILAI = A")
	default:
		fmt.Println("NILAI TIDAK DI KETAHUI")
	}

	switch {
	case x == 60:
		fmt.Println("NILAI = C")
	case x == 80:
		fmt.Println("NILAI = B")
	case x == 100:
		fmt.Println("NILAI = A")
	default:
		fmt.Println("NILAI TIDAK DI KETAHUI")

	}

	// type switch

	var y interface{}
	y = "Wuriyanto"

	switch y.(type) {
	case int:
		fmt.Println("y bertipe data int")
	case string:
		fmt.Println("y bertipe data string")
	case float64:
		fmt.Println("y bertipe data float64")
	default:
		fmt.Println("y tipe datanya tidak di ketahui")
	}

}
