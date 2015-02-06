// +build ignore

package main

import (
	"compress/gzip"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	for _, arg := range flag.Args() {
		func() {
			// START1 OMIT
			fh, err := os.Open(arg)
			if err != nil {
				log.Fatal(err)
			}
			defer fh.Close()

			gz, err := gzip.NewReader(fh)
			if err != nil {
				log.Fatal(err)
			}
			defer gz.Close()

			buf, err := ioutil.ReadAll(gz)
			if err != nil {
				log.Fatal(err)
			}

			var mine MyStruct
			err = json.Unmarshal(&mine)
			if err != nil {
				log.Fatal(err)
			}
			// STOP1 OMIT

		}()
	}

}

type MyStruct struct {
}
