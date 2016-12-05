package main

import (
  "log"
  "net/http"
  _ "net/http/pprof"
  "time"
)

func main() {
  go func(){ log.Println(http.ListenAndServe("localhost:6060",nil))}()
  log.Println("Sleeping")
  for {
    log.Print(".")
    time.Sleep(3 * time.Second)
  }
}

