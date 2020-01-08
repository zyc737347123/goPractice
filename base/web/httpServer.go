package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/zyc737347123/goPractice/base/funcs"
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

// PrintRequest will response request info
func PrintRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// GetGif will response a gif
func GetGif(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}

	funcs.Lissajous1(w)

	if len(r.Form["cycles"]) > 0 {
		fmt.Println(r.Form["cycles"][0])
		cycles, _ := strconv.ParseFloat(r.Form["cycles"][0], 64)
		funcs.Lissajous2(w, cycles)
	} else {
		funcs.Lissajous1(w)
	}
}
