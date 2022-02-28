package main

import (
	"fmt"
	"gomodtest/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(4)

	root.Traverse()
	fmt.Println("----------------")
	root.TraverseFunc(func(node *tree.Node) {
		node.Print()
	})

	fmt.Println("----------------")
	count := 0
	root.TraverseFunc(func(node *tree.Node) {
		count++
	})
	fmt.Println(count)
}
