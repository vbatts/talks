// +build OMIT
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

// START1 OMIT
type MyStruct struct {
	Name     string    `json:"myname"`
	DateTime time.Time `json:"mytime"`
}

// STOP1 OMIT

func main() {
	var (
		buf []byte
		err error
	)
	// START2 OMIT
	m := MyStruct{"vbatts", time.Now()}
	fmt.Printf("%#v\n", m)
	if buf, err = xml.Marshal(m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(buf))

	var m1 MyStruct
	if err = xml.Unmarshal(buf, &m1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", m1)

	// STOP2 OMIT
}
