package main

import (
	"log"
	"net/http"

	"github.com/zyc737347123/goPractice/base/web"
)

func main() {
	// your http.Handle calls here
	http.Handle("/string", web.String("I'm a link zhang."))
	http.Handle("/struct", &web.Struct{Who: "Hello", Punct: ":", Greeting: "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
