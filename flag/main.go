package main

import (
  "fmt"
  "flag"
)

func main() {

  username := flag.String("U", "", "Your Username")
  password := flag.String("P", "", "Your Password")

  flag.Parse()

  result := login(*username, *password)

  if result {
    fmt.Println("Login Success")
  } else {
    fmt.Println("Login Fail")
  }
}

func login(username, password string) bool {
  if username == "Wury" && password == "123456" {
    return true
  }
  return false
}
