package main

import (
	"fmt"
	"os"

	"github.com/zyc737347123/goPractice/base/crawler"
	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	crawler.Title(os.Args[1])
	_, n, err := funcs.Fetch2(os.Args[1])
	fmt.Println(n, err)
}
