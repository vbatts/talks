// +build ignore

package main

// START1 OMIT
import (
	"fmt"
	"os"

	"github.com/foo/bar"
)

// STOP1 OMIT

func main() {
	fmt.Println(bar.Baz())
	_ = os.FileInfo
}
