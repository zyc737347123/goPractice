// Fetch prints the content found at a URL.
package main

import (
	"os"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	for _, url := range os.Args[1:] {
		res := funcs.Fetch(url)
		if !res {
			os.Exit(1)
		}
	}
}
