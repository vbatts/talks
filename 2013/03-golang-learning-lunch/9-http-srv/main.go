package main

import (
  "fmt"
  "log"
  "net/http"
)

// START OMIT
func routeSlash(w http.ResponseWriter, r *http.Request) {
  log.Printf("Got a request from %s", r.RemoteAddr)
  fmt.Fprintf(w, "It Works!\n")
}

func main() {
  log.Println("listening on 0.0.0.0:4000")
  http.HandleFunc("/", routeSlash)
  log.Fatal(http.ListenAndServe("0.0.0.0:4000", nil))
}
// STOP OMIT

