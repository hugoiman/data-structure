package SingleLinkedList

import (
	"fmt"
	"os"
)

type SingleLinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

func Init() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (sll *SingleLinkedList) IsEmpty() bool {
	return (sll.Head == SingleLinkedList{}.Head)
}

func (sll *SingleLinkedList) GetSize() int {
	return sll.Size
}

func (sll *SingleLinkedList) DisplayData() {
	temp := sll.Head
	fmt.Print("Data List: ")
	for temp != nil {
		fmt.Printf("%d, ", temp.Data)
		temp = temp.Next
	}
	fmt.Println(fmt.Sprintf("Head: %d, Tail: %d, Size: %d ", sll.Head.Data, sll.Tail.Data, sll.Size))
}

func (sll *SingleLinkedList) AddFirst(input int) {
	newNode := new(Node)
	newNode.Node(input)
	if sll.IsEmpty() {
		sll.Head = newNode
		sll.Tail = newNode
	} else {
		newNode.Next = sll.Head
		sll.Head = newNode
	}
	sll.Size++
}

func (sll *SingleLinkedList) AddLast(input int) {
	newNode := new(Node)
	newNode.Node(input)
	if sll.IsEmpty() {
		sll.Head = newNode
		sll.Tail = newNode
	} else {
		sll.Tail.Next = newNode
		sll.Tail = newNode
	}
	sll.Size++
}

func (sll *SingleLinkedList) InsertAfter(input, data int) {
	newNode := new(Node)
	newNode.Node(input)
	temp := sll.Head

	for temp != nil {
		if temp.Data == data {
			newNode.Next = temp.Next
			temp.Next = newNode
			sll.Size++
			break
		} else if temp == sll.Tail && temp.Data != data {
			fmt.Printf("Data %d not found. ", data)
		}
		temp = temp.Next
	}
}

func (sll *SingleLinkedList) RemoveFirst() {
	temp := sll.Head
	if !sll.IsEmpty() {
		if sll.Head == sll.Tail {
			sll.Head, sll.Tail = nil, nil
		} else {
			sll.Head = temp.Next
		}
		sll.Size--
	} else {
		fmt.Print("Data is empty")
	}
}

func (sll *SingleLinkedList) RemoveLast() {
	temp := sll.Head
	if !sll.IsEmpty() {
		if sll.Tail == sll.Head {
			sll.Head, sll.Tail = nil, nil
		} else {
			for temp.Next != sll.Tail {
				temp = temp.Next
			}
			temp.Next = nil
			sll.Tail = temp
		}
		sll.Size--
	} else {
		fmt.Print("Data is empty")
	}
}

func (sll *SingleLinkedList) RemoveCertain(data int) {
	temp := sll.Head
	if !sll.IsEmpty() {
		for temp != nil {
			if temp == sll.Tail && temp.Data != data {
				fmt.Printf("Data %d not found. ", data)
				break
			} else if temp.Next.Data == data {
				temp.Next = temp.Next.Next
				if temp.Next == nil {
					sll.Tail = temp
				}
				sll.Size--
				break
			} else if temp.Data == data && temp == sll.Head {
				sll.RemoveFirst()
				sll.Size--
				break
			}
			temp = temp.Next
		}
	} else {
		fmt.Print("Data is empty")
	}
}

func (sll *SingleLinkedList) CheckIndex(index int) {
	if index < 0 || index > sll.Size {
		fmt.Printf("index %d out of bounds exception ", index)
		os.Exit(0)
	}
}

func (sll *SingleLinkedList) GetIndex(index int) int {
	sll.CheckIndex(index)
	temp := sll.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp.Data
}
