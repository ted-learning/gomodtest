package tree

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
