package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var myString = "Hello Golang"
	var myStringTwo = "10"

	fmt.Println(myString)
	fmt.Println(len(myString))

	myStringUpper := strings.ToUpper(myString)
	fmt.Println(myStringUpper)

	myStringLower := strings.ToLower(myString)
	fmt.Println(myStringLower)

	resultContains := strings.Contains(myString, "go") // Go != go
	fmt.Println(resultContains)

	resultSplit := strings.Split(myString, " ") // return dynamyc array/ slice
	for _, v := range resultSplit {
		fmt.Println(v)
	}

	// cara mengkonversi string menjadi integer
	resultConvInt, err := strconv.Atoi(myStringTwo)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultConvInt * 5)

	// cara mengkonversi integer menjadi string

	resultConvStr := strconv.Itoa(10)
	fmt.Println(resultConvStr)
}
