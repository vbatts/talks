package main

import (
  "fmt"
  "time"
)

type Person struct {
  Name string
  Dob *time.Time
}

func main() {
  // START OMIT
  t, err := time.Parse(time.RFC822, "27 Mar 75 00:00 EST")
  if (err != nil)  {
    fmt.Println(err)
    return
  }
  p := Person{Name: "John Doe", Dob: &t }
  fmt.Printf("%q\n", p)
  // STOP OMIT

  //fmt.Printf("%s\n", p.Name)
  //fmt.Printf("%s\n", p.Dob.String())
}

