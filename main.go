package main

import (
	"fmt"
	"os"

	"github.com/zyc737347123/goPractice/base/crawler"
)

func main() {
	fmt.Println(crawler.CountWordsAndImages(os.Args[1]))
}
