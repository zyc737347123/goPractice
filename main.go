// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"os"

	"github.com/zyc737347123/goPractice/base/funcs"
)

// LineInfo is a struct for record text line info
type LineInfo struct {
	name string
	text string
}

func main() {
	files := os.Args[1:]
	funcs.CountFilesLines(files)
}

func countLines(f *os.File, counts map[LineInfo]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		lineInfo := LineInfo{name: fileName, text: input.Text()}
		counts[lineInfo]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
