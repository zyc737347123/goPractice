package main

import (
	"fmt"
	"os"

	"github.com/zyc737347123/goPractice/base/crawler"
)

func main() {
	doc, err := crawler.FindLinks2(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	res := crawler.ElementByTagName(doc, (os.Args[2:])...)
	for _, n := range res {
		fmt.Println(n.Data)
		for _, v := range n.Attr {
			fmt.Printf(" %s='%s'", v.Key, v.Val)
		}
		fmt.Println()
	}
}
