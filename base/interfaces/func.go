package interfaces

import (
	"bufio"
	"fmt"
	"io"
)

// An WordCounter can count word and line num
type WordCounter struct {
	words int
	lines int
}

// An ByteCounter can count word
type ByteCounter struct {
	w     io.Writer
	count int64
}

func (w *WordCounter) Write(p []byte) (int, error) {
	for index := 0; index < len(p); {
		i, v, err := bufio.ScanWords(p[index:], true)
		if err != nil {
			return index, err
		}
		index += i
		w.words++
		fmt.Printf("word: %s\n", v)
	}

	for index := 0; index < len(p); {
		i, v, err := bufio.ScanLines(p[index:], true)
		if err != nil {
			return index, err
		}
		index += i
		w.lines++
		fmt.Printf("link %s\n", v)
	}
	return len(p), nil
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	l, err := b.w.Write(p)
	b.count += int64(l)
	return l, err
}

// CountingWriter will return pack io writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	b := new(ByteCounter)
	b.w = w
	return b, &b.count
}
