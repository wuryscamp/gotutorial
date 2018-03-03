package main

import (
	"fmt"
)

func main() {

	lv := love("Wuriyanto")
	fmt.Println(lv("Golang"))
	fmt.Println(lv("Java"))

}

func genNumber() func() int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}
}

func love(name string) func(string) string {
	return func(things string) string {
		return fmt.Sprintf("%s love %s", name, things)
	}
}
