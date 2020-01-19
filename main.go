package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zyc737347123/goPractice/base/crawler"
)

func main() {
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	if err := crawler.WaitForServer(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		log.Fatalf("Site is down: %v\n", err)
		os.Exit(1)
	}
}
