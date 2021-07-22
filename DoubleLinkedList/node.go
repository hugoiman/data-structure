package DoubleLinkedList

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func (n *Node) Node(data int) {
	n.Data = data
	n.Next = nil
	n.Prev = nil
}
