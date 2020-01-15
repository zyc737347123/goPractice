// Server2 is a minimal "echo" and counter server.
package main

import (
	"log"
	"os"

	"github.com/zyc737347123/goPractice/base/text"
	"github.com/zyc737347123/goPractice/base/web"
)

func main() {
	result, err := web.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	text.Report1(result)
}
