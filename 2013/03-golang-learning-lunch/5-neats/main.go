package main

import "fmt"

func main() {
  numbers := []int{1,7,3,8,2,4,5,4,9,0}

  for i := range numbers {
    fmt.Printf("index: %d\n", i)
  }

  for _, v := range numbers {
    fmt.Printf("value: %d\n", v)
  }

  people := map[string]Person{}
  people["shawking"] = Person{ Name: "Stephen Hawking", State: LIVING}
  people["dritchie"] = Person{ Name: "Dennis Ritchie", State: DEAD}
  for k,v := range people {
    fmt.Printf("Key: %s; Object: %q\n", k, v)
  }

}

