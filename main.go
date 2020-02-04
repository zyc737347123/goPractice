package main

import (
	"os"

	"github.com/zyc737347123/goPractice/base/crawler"
	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	funcs.BFS(crawler.Crawl, os.Args[1:])
}
