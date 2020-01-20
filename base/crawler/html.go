package crawler

import (
	"fmt"

	"golang.org/x/net/html"
)

var depth int
var Id string

// ElementByID is a func
func ElementByID(doc *html.Node, id string) *html.Node {
	Id = id
	return forEachNode(doc, findByID)
}

func findByID(n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, v := range n.Attr {
			if v.Key == "id" && v.Val == Id {
				return true
			}
		}
	}
	return false
}

func forEachNode(n *html.Node, take func(*html.Node) bool) *html.Node {
	if take != nil && take(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if res := forEachNode(c, take); res != nil {
			return res
		}
	}

	return nil
}

// ForEachNode is a func, traversal html doc
func ForEachNode(n *html.Node, pre, post func(*html.Node) bool) {

	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// StartElement is a func
func StartElement(n *html.Node) bool {
	if n.Type == html.ElementNode && n.FirstChild == nil {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		printAttr(n)
	} else if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		printAttr(n)
		fmt.Printf(">\n")
		depth++
	}

	if n.Type == html.TextNode || n.Type == html.CommentNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}

	return true
}

// EndElement is a func
func EndElement(n *html.Node) bool {
	if n.Type == html.ElementNode && n.FirstChild == nil {
		fmt.Printf("/>\n")
	} else if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return true
}

func printAttr(n *html.Node) {
	for _, v := range n.Attr {
		fmt.Printf(" %s='%s'", v.Key, v.Val)
	}
}
