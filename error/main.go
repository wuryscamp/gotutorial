package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	iStr := "a"

	i, err := strconv.Atoi(iStr)

	if err != nil {
		fmt.Println("Terjadi error : ", err.Error())
	} else {
		fmt.Println(i)
	}

	r, err := Div(10, 0)

	if err != nil {
		fmt.Println("Terjadi error : ", err.Error())
	} else {
		fmt.Println(r)
	}
	// inline handling error
	if x, err := Div(25, 0); err != nil {
		fmt.Println("Terjadi error : ", err.Error())
	} else {
		fmt.Println(x)
	}
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak Bisa Membagi dengan 0")
	}
	result := x / y
	return result, nil
}
