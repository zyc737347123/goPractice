package main

import (
	"log"
	"net/http"

	"github.com/zyc737347123/goPractice/base/web"
)

func main() {
	var h web.Hello
	err := http.ListenAndServe("localhost:5000", h)
	if err != nil {
		log.Fatal(err)
	}
}
