package funcs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
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

// FetchAll will fetch all url and show time
func FetchAll(urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
}

// fetch use channel output result
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// Fetch2 output result to file
func Fetch2(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	fmt.Println(local)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if err == nil {
			err = f.Close()
			fmt.Println(err)
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err

}
