package funcs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Fetch will fetch url and print to stdout
func Fetch(url string) bool {
	if !strings.HasPrefix(url, "http://") {
		fmt.Fprintf(os.Stderr, "not supprot prefix\n")
		return false
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return false
	}
	n, err := io.Copy(os.Stdout, resp.Body)
	fmt.Printf("request status %d, get %d byte\n", resp.StatusCode, n)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return false
	}
	return true
}
