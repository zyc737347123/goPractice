package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// LineInfo is a struct for record text line info
type LineInfo struct {
	name string
	text string
}

func countLines(f *os.File, counts map[LineInfo]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		lineInfo := LineInfo{name: fileName, text: input.Text()}
		counts[lineInfo]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

// CountFilesLines will count all files same line
func CountFilesLines(files []string) {
	counts := make(map[LineInfo]int)
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s: %d\t%s\n", line.name, n, line.text)
		}
	}
}

func max(val int, vals ...int) int {
	res := val
	for _, v := range vals {
		if res < v {
			res = v
		}
	}
	return res
}

func min(val int, vals ...int) int {
	res := val
	for _, v := range vals {
		if res > v {
			res = v
		}
	}
	return res
}
