package SingleLinkedList

import (
	"errors"
	"fmt"
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
	return (sll.Size == 0)
}

func (sll *SingleLinkedList) GetSize() int {
	return sll.Size
}

func (sll *SingleLinkedList) DisplayData() {
	if !sll.IsEmpty() {
		temp := sll.Head
		fmt.Print("Data List: ")
		for temp != nil {
			fmt.Printf("%d, ", temp.Data)
			temp = temp.Next
		}
		fmt.Println(fmt.Sprintf("Head: %d, Tail: %d, Size: %d ", sll.Head.Data, sll.Tail.Data, sll.Size))
	} else {
		fmt.Println("Data is empty")
	}
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

func (sll *SingleLinkedList) InsertAfter(input, dataAfter int) {
	newNode := new(Node)
	newNode.Node(input)
	temp := sll.Head

	for temp != nil {
		if temp.Data == dataAfter {
			newNode.Next = temp.Next
			temp.Next = newNode
			sll.Size++
			break
		} else if temp == sll.Tail && temp.Data != dataAfter {
			fmt.Printf("Data %d not found. ", dataAfter)
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

func (sll *SingleLinkedList) CheckIndex(index int) error {
	if index < 0 || index > sll.Size {
		return errors.New("index out of bounds exception")
	}
	return nil
}

func (sll *SingleLinkedList) GetIndex(index int) interface{} {
	err := sll.CheckIndex(index)
	if err != nil {
		return err
	}
	temp := sll.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp.Data
}
