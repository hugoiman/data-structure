package DoubleLinkedList

import (
	"errors"
	"fmt"
)

type DoubleLinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

func Init() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (dll *DoubleLinkedList) IsEmpty() bool {
	return (dll.Size == 0)
}

func (dll *DoubleLinkedList) GetSize() int {
	return dll.Size
}

func (dll *DoubleLinkedList) DisplayData() {
	if !dll.IsEmpty() {
		temp := dll.Head
		fmt.Print("Data List: ")
		for temp != nil {
			fmt.Printf("%d, ", temp.Data)
			temp = temp.Next
		}
		fmt.Println(fmt.Sprintf("Head: %d, Tail: %d, Size: %d ", dll.Head.Data, dll.Tail.Data, dll.Size))
	} else {
		fmt.Println("Data is empty")
	}
}

func (dll *DoubleLinkedList) DisplayDataFromBack() {
	if !dll.IsEmpty() {
		temp := dll.Tail
		fmt.Print("Data List: ")
		for temp != nil {
			fmt.Printf("%d, ", temp.Data)
			temp = temp.Prev
		}
		fmt.Println(fmt.Sprintf("Head: %d, Tail: %d, Size: %d ", dll.Head.Data, dll.Tail.Data, dll.Size))
	} else {
		fmt.Println("Data is empty")
	}
}

func (dll *DoubleLinkedList) AddFirst(input int) {
	newNode := new(Node)
	newNode.Node(input)
	if dll.IsEmpty() {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
	dll.Size++
}

func (dll *DoubleLinkedList) AddLast(input int) {
	newNode := new(Node)
	newNode.Node(input)
	if dll.IsEmpty() {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}
	dll.Size++
}

func (dll *DoubleLinkedList) InsertAfter(input, dataAfter int) {
	newNode := new(Node)
	newNode.Node(input)
	temp := dll.Head

	for temp != nil {
		if temp.Data == dataAfter {
			newNode.Next = temp.Next
			newNode.Prev = temp
			temp.Next.Prev = newNode
			temp.Next = newNode
			dll.Size++
			break
		} else if temp == dll.Tail && temp.Data != dataAfter {
			fmt.Printf("Data %d not found. ", dataAfter)
		}
		temp = temp.Next
	}
}

func (dll *DoubleLinkedList) InsertBefore(input, dataBefore int) {
	newNode := new(Node)
	newNode.Node(input)
	temp := dll.Head

	for temp != nil {
		if temp.Data == dataBefore {
			if temp == dll.Head {
				dll.AddFirst(input)
			} else {
				newNode.Next = temp
				newNode.Prev = temp.Prev
				temp.Prev.Next = newNode
				temp.Prev = newNode
				dll.Size++
			}
			break
		} else if temp == dll.Tail && temp.Data != dataBefore {
			fmt.Printf("Data %d not found. ", dataBefore)
		}
		temp = temp.Next
	}
}

func (dll *DoubleLinkedList) RemoveFirst() {
	temp := dll.Head
	if !dll.IsEmpty() {
		if dll.Head == dll.Tail {
			dll.Head, dll.Tail = nil, nil
		} else {
			dll.Head.Next.Prev = nil
			dll.Head = temp.Next
		}
		dll.Size--
	} else {
		fmt.Print("Data is empty")
	}
}

func (dll *DoubleLinkedList) RemoveLast() {
	temp := dll.Tail
	if !dll.IsEmpty() {
		if dll.Tail == dll.Head {
			dll.Head, dll.Tail = nil, nil
		} else {
			dll.Tail.Prev.Next = nil
			dll.Tail = temp.Prev
		}
		dll.Size--
	} else {
		fmt.Print("Data is empty")
	}
}

func (dll *DoubleLinkedList) RemoveCertain(data int) {
	temp := dll.Head
	if !dll.IsEmpty() {
		for temp != nil {
			if temp == dll.Tail && temp.Data != data {
				fmt.Printf("Data %d not found. ", data)
				break
			} else if temp.Data == data {
				if temp == dll.Head {
					dll.RemoveFirst()
				} else if temp == dll.Tail {
					dll.RemoveLast()
				} else if temp != dll.Head && temp != dll.Tail {
					temp.Prev.Next = temp.Next
					temp.Next.Prev = temp.Prev
					dll.Size--
				}
				break
			}
			temp = temp.Next
		}
	} else {
		fmt.Print("Data is empty")
	}
}

func (dll *DoubleLinkedList) CheckIndex(index int) error {
	if index < 0 || index > dll.Size {
		return errors.New("index out of bounds exception")
	}
	return nil
}

func (dll *DoubleLinkedList) GetIndex(index int) interface{} {
	err := dll.CheckIndex(index)
	if err != nil {
		return err
	}
	temp := dll.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp.Data
}
