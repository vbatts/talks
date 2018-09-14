package main

func main() {
  jeep := Hummer{}
  h(jeep)
  jeep.Where()
  c(jeep)
}


type Car interface {
  Honk()
  Crank()
}

func c(some interface{}) {
  some.(Car).Crank()
}
func h(some interface{}) {
  some.(Car).Honk()
}

// Implement the Car interface
type Hummer struct {
  Car
}

func (h Hummer) Honk() {
  println("BEEP BEEP")
}

func (h Hummer) Where() {
  println("Who's got the keys to the jeep?")
}

func (h Hummer) Crank() {
  println("VROOOOOOM")
}

