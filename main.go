// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zyc737347123/goPractice/base/web"
)

func main() {
	result, err := web.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
