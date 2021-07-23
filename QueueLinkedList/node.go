package QueueLinkedList

type Node struct {
	Data int
	Next *Node
}

func (n *Node) Node(data int) {
	n.Data = data
	n.Next = nil
}
