package graph

import "fmt"

type Node struct {
	url    string
	parent *Node
	child  []*Node
}

func (n *Node) SetUrl(url string) {
	n.url = url

}

func (n Node) GetUrl() string {
	return n.url
}

func (p *Node) Add(url string) bool {
	n := new(Node)
	n.url = url
	n.parent = p
	p.child = append(p.child, n)
	return true
}

func (n Node) GetChild() []*Node {
	return n.child
}

func (n Node) GetParent() Node {
	return n
}

func (n Node) Print() {
	for i := 0; i < len(n.child); i++ {
		fmt.Println(n.child[i].GetUrl())
		if n.child[i].child != nil {
			n.child[i].Print()
		}
	}
}
