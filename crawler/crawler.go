package main

import (
	"fmt"
	gr "github.com/blackdev1l/go-crawler/graph"
	ht "golang.org/x/net/html"
	"net/http"
)

func main() {
	tree := gr.Node{}
	tree.SetUrl("google.com")
	url := tree.GetUrl()
	fmt.Println(url)
	tree.Add("yahoo.com")
	tree.Add("mage.it")
	tree.Add("leroy.com")
	child := tree.GetChild()
	child[2].Add("mage.com")
	p := child[2].GetParent()
	tree.Print()
	resp, err := http.Get("http://gaming.ngi.it/index.php")
	if err != nil {
		p = ht.Parse(resp.Body)
		fmt.Println(p)
	}
	p.Print()
}
