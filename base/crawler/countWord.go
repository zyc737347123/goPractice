package crawler

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {

	texts, images := visit4(nil, 0, n)
	for _, v := range texts {
		v = strings.Trim(strings.TrimSpace(v), "\r\n")
		if v == "" {
			continue
		}
		words += strings.Count(v, "")
	}
	return
}

func visit4(texts []string, imgs int, n *html.Node) ([]string, int) {
	//文本
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	//图片
	if n.Type == html.ElementNode && (n.Data == "img") {
		imgs++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}

		texts, imgs = visit4(texts, imgs, c)
	}
	//多返回值
	return texts, imgs
}
