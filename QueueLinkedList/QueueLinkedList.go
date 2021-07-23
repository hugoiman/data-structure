package QueueLinkedList

import "fmt"

type QueueLinkedList struct {
	Head  *Node
	Tail  *Node
	Front *Node
	Rear  *Node
}

func Init() *QueueLinkedList {
	return &QueueLinkedList{}
}

func (qll *QueueLinkedList) IsEmpty() bool {
	return (qll.Rear == nil)
}

func (qll *QueueLinkedList) DisplayData() {
	if !qll.IsEmpty() {
		temp := qll.Head
		fmt.Print("Data List: ")
		for temp != nil {
			fmt.Printf("<- %d ", temp.Data)
			temp = temp.Next
		}
		fmt.Println(fmt.Sprintf(". Head/Front: %d, Tail/Rear: %d ", qll.Head.Data, qll.Tail.Data))
	} else {
		fmt.Println("Data is empty")
	}
}

func (qll *QueueLinkedList) AddLast(input int) {
	newNode := new(Node)
	newNode.Node(input)
	if qll.IsEmpty() {
		qll.Head = newNode
		qll.Tail = newNode
		qll.Front = qll.Head
		qll.Rear = qll.Tail
	} else {
		qll.Tail.Next = newNode
		qll.Tail = newNode
	}
}

func (qll *QueueLinkedList) RemoveFirst() {
	temp := qll.Head
	if qll.Head == qll.Tail {
		qll.Head = nil
		qll.Tail = nil
		qll.Rear = nil
	} else {
		temp = temp.Next
		qll.Head = temp
		temp = nil
	}
}

func (qll *QueueLinkedList) Enqueue(input int) {
	qll.AddLast(input)
	qll.Rear = qll.Tail
}

func (qll *QueueLinkedList) Dequeue() int {
	if !qll.IsEmpty() {
		temp := qll.Front.Data
		qll.RemoveFirst()
		qll.Front = qll.Head
		return temp
	} else {
		fmt.Println("Data is empty")
		return -1
	}
}

func (qll *QueueLinkedList) Peek() int {
	return qll.Front.Data
}
