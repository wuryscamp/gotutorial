package main

import (
  "fmt"
  "time"
)

func main() {

  t := time.Now()

  fmt.Println(t)
  fmt.Println(t.Year())
  fmt.Println(t.Month())
  fmt.Println(t.Day())

  fmt.Println("====================")

  timeString := "January 20, 2018"
  form := "January 2, 2006"

  resTime, err := time.Parse(form, timeString)

  if err != nil{
    fmt.Println(err)
  }

  fmt.Println(resTime)
  fmt.Println(resTime.Year())

  t1 := time.Date(2015, 2, 1, 12, 0, 0, 0, time.UTC)
  t2 := time.Date(2015, 2, 1, 12, 0, 0, 0, time.UTC)

  eq := t1.Equal(t2)

  fmt.Println(t1)
  fmt.Println(eq)

  layout := "2006-02-01"

  timeNow := time.Now()

  resString := timeNow.Format(layout)

  fmt.Println(resString)

}
