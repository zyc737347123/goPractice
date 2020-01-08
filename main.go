// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"os"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	files := os.Args[1:]
	funcs.CountFilesLines(files)
}
