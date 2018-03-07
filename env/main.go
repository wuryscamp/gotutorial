package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("POSTGRES", "postgres.two.host")

	postgres := os.Getenv("POSTGRES")
	mysql := os.Getenv("MYSQL")

	fmt.Println("Postgres Config : ", postgres)
	fmt.Println("MySQL Config : ", mysql)
}
