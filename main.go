package main

import (
	"os"

	"github.com/zyc737347123/goPractice/base/crawler"
)

func main() {
	doc, _ := crawler.FindLinks2(os.Args[1])
	crawler.ForEachNode(doc, crawler.StartElement, crawler.EndElement)
}
