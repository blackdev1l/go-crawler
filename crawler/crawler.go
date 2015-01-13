package main

import (
	"fmt"
	gr "github.com/blackdev1l/go-crawler/graph"
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
}
