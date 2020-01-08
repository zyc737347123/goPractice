// Server2 is a minimal "echo" and counter server.
package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/zyc737347123/goPractice/base/web"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", web.PrintRequest)
	http.HandleFunc("/gif", web.GetGif)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
