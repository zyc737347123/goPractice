package crawler

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type visitF func(links []string, n *html.Node) []string

// Visit is func
func Visit(file *os.File, option string) {
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	var v visitF = visit

	switch option {
	case "visit":
		v = visit
	case "visitR":
		v = visitR
	case "visit3":
		v = visit3
	}

	for i, link := range v(nil, doc) {
		fmt.Println(i, link)
	}

	fmt.Println("finish")
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
			fmt.Println(a)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func visitR(links []string, n *html.Node) []string {

	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visitR(links, n.FirstChild)
	links = visitR(links, n.NextSibling)

	return links
}

func count(res map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		res[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = count(res, c)
	}
	return res
}

// Count is a func
func Count(file *os.File) {
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	p := make(map[string]int)

	for k, v := range count(p, doc) {
		fmt.Println(k, v)
	}
}

func visit3(texts []string, n *html.Node) []string {

	if n.Data == "script" || n.Data == "style" {
		return texts
	}

	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = visit3(texts, c)
	}
	return texts
}
