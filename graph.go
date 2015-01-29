package main

import (
	"encoding/json"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strconv"
)

type Edge struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type Graph struct {
	Url   string  `json:"-"`
	Child []*Node `json:"nodes,omitempty"`
	Edges []*Edge `json:"edges,omitempty"`
}
type Node struct {
	Url   string `json:"id,omitempty"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Size  int    `json:"size"`
	Label string `json:"label,omitempty"`
	//	Child   []*Node `json:"nodes,omitempty"`
	//	Edges   []*Edge `json:"edges,omitempty"`
}

func (p *Graph) Add(url string) bool {
	n := new(Node)
	e := new(Edge)
	n.Label = url
	num := len(p.Child)

	log.Println(strconv.Itoa(num))
	n.Url = strconv.Itoa(len(p.Child))
	e.Id = "e-" + strconv.Itoa(len(p.Child))
	e.Source = strconv.Itoa(len(p.Child))
	e.Target = "0"
	//setUrl(p.Url, url, n, e)

	if (len(p.Child))%2 != 0 {
		n.X = p.Child[len(p.Child)-1].X
		n.X += 5
	} else {

		n.X = p.Child[len(p.Child)-2].X
		n.X -= 5
	}
	n.Size = 1
	//n.parent = p
	if url != p.Url {
		p.Edges = append(p.Edges, e)
	}
	p.Child = append(p.Child, n)
	return true
}

func (P *Graph) Parse(url string, number int) []byte {
	count := 0
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node, *Graph, int, *int)
	f = func(n *html.Node, P *Graph, number int, count *int) {
		if n.Type == html.ElementNode && n.Data == "a" && *count < number {
			P.Add(n.Attr[0].Val)
			*count += 1
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, P, number, count)
		}
	}
	f(doc, P, number, &count)

	b, err := json.Marshal(P)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

/*
	Set id and label of both node and edge with the url
*/

func setUrl(target string, url string, n *Node, e *Edge) {
	eID := "e-" + url
	e.Id = eID
	e.Id = eID
	e.Source = url
	e.Target = target
	n.Url = url
	n.Label = url
	return
}
