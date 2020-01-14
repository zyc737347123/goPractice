// Server2 is a minimal "echo" and counter server.
package main

import (
	"os"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	funcs.Wordfreq(os.Args[1])
}
