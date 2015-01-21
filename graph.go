package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

type Edge struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type Node struct {
	Url    string `json:"id,omitempty"`
	X      int    `json:"x,omitempty"`
	Y      int    `json:"y,omitempty"`
	Size   int    `json:"size,omitempty"`
	Label  string `json:"label,omitempty"`
	parent *Node
	Child  []*Node `json:"nodes,omitempty"`
	Edges  []*Edge `json:"edges,omitempty"`
}

func NewGraph() *Node {
	n := new(Node)
	n.X = 2
	n.Y = 1
	return n
}

func (p *Node) Add(url string) bool {
	n := new(Node)
	e := new(Edge)
	eID := "e-" + url
	e.Id = eID
	e.Id = eID
	e.Source = url
	e.Target = p.Url
	n.Url = url
	n.Label = url
	n.X += len(p.Child) + 1
	n.Y += 1
	n.Size = 1
	//n.parent = p
	if url != p.Url {
		p.Edges = append(p.Edges, e)
	}
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
		log.Fatal(err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	return b
}
