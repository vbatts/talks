package main

func main() {
  // START OMIT
  words := []string{"cow","goat","sheep","horse","chicken"}
  for i := range words {
    println(words[i])
  }
  for _, word := range words {
    println(word)
  }
  // STOP OMIT
}

