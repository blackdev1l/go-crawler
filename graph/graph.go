package graph

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

type Node struct {
	Url    string `json:"url"`
	parent *Node
	Child  []*Node
}

func (p *Node) Add(url string) bool {
	n := new(Node)
	n.Url = url
	//n.parent = p
	p.Child = append(p.Child, n)
	return true
}

func (n Node) Print() {
	for i := 0; i < len(n.Child); i++ {
		fmt.Println(n.Child[i].Url)
		if n.Child[i].Child != nil {
			n.Child[i].Print()
		}
	}
}

func (P *Node) Parse(url string) []byte {
	P.Url = url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error at http.get ")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("error at html.parse ")
	}

	var f func(*html.Node, *Node)
	f = func(n *html.Node, P *Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			P.Add(n.Attr[0].Val)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, P)
		}
	}
	f(doc, P)
	b, err := json.Marshal(P)
	if err != nil {
		fmt.Println(err.Error())
	}
	return b

}
