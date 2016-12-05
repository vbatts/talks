package main

import (
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

