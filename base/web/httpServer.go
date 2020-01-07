package web

import (
	"fmt"
	"net/http"
)

// Hello is a web handle struct
type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// String is a string web handle
type String string

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, str)
}

// Struct is a struct web handle
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "%s%s%s\n", str.Who, str.Punct, str.Greeting)
}
