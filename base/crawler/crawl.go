package crawler

import "fmt"

// Crawl will crawl url
func Crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		fmt.Println("extract url fail", err)
	}
	return list
}
