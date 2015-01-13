package main

import (
	"fmt"
	gr "github.com/blackdev1l/go-crawler/graph"
	"golang.org/x/net/html"
	"net/http"
)

func main() {

}

func parse(url string) {
	resp, err := http.Get("http://gaming.ngi.it/index.php")
	if err != nil {
		fmt.Println("error at http.get ")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("error at html.parse ")
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			fmt.Println(n.Attr[0].Val)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
