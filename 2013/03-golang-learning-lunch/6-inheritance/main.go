
package main

import (
  "log"
  "os"
  "os/exec"
  "time"
)

type Job struct {
  Command string
  *log.Logger
}

func NewJob(cmd string) *Job {
  return &Job{ cmd, log.New(os.Stderr, "[Job] ", log.Ldate)}
}

func (j *Job) Run() ([]byte, error) {
  j.Printf("Running [%s] ...", j.Command)
  out, err := exec.Command(j.Command).Output()
  if (err != nil) {
    j.Fatal(err)
    return nil, err
  }
  j.Printf("Command returned [%s]", out)
  return out, nil
}

func main() {
  job := NewJob("uptime")
  job.Run()

  for i := 0; i < 10; i++ {
    job.Run()
    time.Sleep(1 * time.Second)
  }
}

