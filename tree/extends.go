package tree

import "fmt"

func (node *Node) Print() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() <-chan *Node {
	var c = make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			c <- node
		})
		close(c)
	}()
	return c
}
