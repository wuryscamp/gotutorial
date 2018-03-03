package main

import (
	"fmt"
)

func main() {

	x := 11
	y := 5

	//arithmetic operator

	//penambahan / additional
	zAdd := x + y
	fmt.Println(zAdd)

	//penguranan / substracts
	zSub := x - y
	fmt.Println(zSub)

	// perkalian / multiplies
	zMul := x * y
	fmt.Println(zMul)

	// pembagian / divides
	zDiv := x / y
	fmt.Println(zDiv)

	// sisa hasil bagi/ modulus
	zMod := x % y
	fmt.Println(zMod)

	// relational operator

	i := 10
	j := 5

	fmt.Println(i == j) //false
	fmt.Println(i != j) // true
	fmt.Println(i > j)  //true
	fmt.Println(i < j)  // false
	fmt.Println(i >= j) // true
	fmt.Println(i <= j) // false
}
