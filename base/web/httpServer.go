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
