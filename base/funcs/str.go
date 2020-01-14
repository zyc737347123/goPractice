package funcs

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Comma1 inserts commas in a non-negative decimal integer string.
func Comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma1(s[:n-3]) + "," + s[n-3:]
}

// Comma2 inserts commas in a non-negative decimal integer string.
func Comma2(s string) string {
	if strings.Contains(s, ".") {
		index := strings.IndexByte(s, '.')
		return comma(s[:index]) + s[index:]
	}

	return comma(s)
}

// CompareByte is a func
func CompareByte(s1, s2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			m1[byte(s1[i])]++
			m2[byte(s2[i])]++
		}

		for i, v := range m1 {
			if v != m2[i] {
				return false
			}
		}

		return true
	}

	return false
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)

	if n <= 3 {
		return s
	}

	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')

		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

// Wordfreq will count all word freq
func Wordfreq(filename string) {
	m := make(map[string]int)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		m[input.Text()]++
	}

	fmt.Println(m)
}
