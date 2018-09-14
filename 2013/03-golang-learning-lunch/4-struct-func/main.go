package main

import (
  "fmt"
  "time"
)

const (
  IN_UTERO = iota
  NEW_BORN
  GROWING
  LIVING
  DYING
  DEAD
)

type Person struct {
  Name string
  Dob *time.Time
  State byte
}

func (p *Person) DobString() string {
  if (p.Dob == nil) {
    return ""
  }
  return p.Dob.String()
}

func (p *Person) ComeToLife() {
  t := time.Now()
  p.Dob = &t
  p.State = NEW_BORN
}


func main() {
  // START OMIT
  p := Person{Name: "John Doe", State: IN_UTERO}
  fmt.Printf("%s, %s\n", p.Name, p.DobString())

  p.ComeToLife()
  fmt.Printf("%s, %s\n", p.Name, p.DobString())
  // STOP OMIT
}

