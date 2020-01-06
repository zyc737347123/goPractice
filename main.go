package main

import (
	"io"
	"os"
	"strings"

	"github.com/zyc737347123/goPractice/base/structs"
)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := structs.Rot13Reader{R: s}
	io.Copy(os.Stdout, &r)
}
