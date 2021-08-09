package main

type Node struct {
	Value string
	Left *Node
	Right *Node
}

func NewNode(data string) *Node {
	return &Node{data, nil, nil}
}

func main()  {
	root := NewNode("A")
	left := NewNode("B")
	right := NewNode("C")
	root.Left = left
	root.Right = right
}