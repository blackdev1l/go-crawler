package main

import (
	"fmt"
	gr "github.com/blackdev1l/go-crawler/graph"
)

func main() {
	tree := gr.Node{}
	b := tree.Parse("https://www.polymer-project.org/docs/elements/paper-elements.html#paper-input")
	fmt.Println(b)
}
