package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/zyc737347123/goPractice/base/interfaces"
)

func main() {
	s := "1234567890"
	data := make([]byte, 1024)
	sr := strings.NewReader(s)
	nr := interfaces.LimitReader(sr, 3)
	for n, err := nr.Read(data); err == nil; {
		fmt.Println(string(data[0:n]))
		n, err = nr.Read(data)
		fmt.Println(err)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	var buf bytes.Buffer
	for n, err := f.Read(data); err == nil; {
		buf.Write(data[0:n])
		n, err = f.Read(data)
	}
	nsr := interfaces.NewReader(buf.String())
	for n, err := nsr.Read(data); err == nil; {
		fmt.Println(string(data[0:n]))
		n, err = nsr.Read(data)
		fmt.Println(err)
	}
}
