package main

import (
  "fmt"
  "time"
)

// START OMIT
func DelayedHello() {
  time.Sleep(3 * time.Second)
  fmt.Printf("%d: Hello!\n", time.Now().Unix())
}

func main() {
  go DelayedHello()
  go DelayedHello()
  go func() {
    time.Sleep(3 * time.Second)
    fmt.Printf("%d: Hello!\n", time.Now().Unix())
  }()

  fmt.Printf("%d: Is there anybody in there?!\n", time.Now().Unix())

  for i := 0; i < 5; i++ {
    fmt.Printf("%d: .\n", time.Now().Unix())
    time.Sleep(time.Second)
  }
}
// STOP OMIT

