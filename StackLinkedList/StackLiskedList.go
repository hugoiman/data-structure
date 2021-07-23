package StackLinkedList

import (
	"errors"
	"fmt"
)

type StackLinkedList struct {
	Top  *Node
	Size int
}

func Init() *StackLinkedList {
	return &StackLinkedList{}
}

func (sll *StackLinkedList) IsEmpty() bool {
	return (sll.Top == nil)
}

func (sll *StackLinkedList) GetSize() int {
	return sll.Size
}

func (sll *StackLinkedList) DisplayData() {
	if !sll.IsEmpty() {
		temp := sll.Top
		fmt.Print("Data List: ")
		for temp != nil {
			fmt.Printf("-> %d ", temp.Data)
			temp = temp.Next
		}
		fmt.Println(fmt.Sprintf(". Top: %d", sll.Top.Data))
	} else {
		fmt.Println("Data is empty")
	}
}

func (sll *StackLinkedList) Push(input int) {
	newNode := new(Node)
	newNode.Node(input)

	if sll.IsEmpty() {
		sll.Top = newNode
	} else {
		newNode.Next = sll.Top
		sll.Top = newNode
	}
	sll.Size++
}

func (sll *StackLinkedList) Pop() interface{} {
	if sll.IsEmpty() {
		return errors.New("Data is empty")
	}
	temp := sll.Top
	sll.Top = sll.Top.Next
	sll.Size--
	return temp.Data
}

func (sll *StackLinkedList) Peek() interface{} {
	if sll.IsEmpty() {
		return errors.New("Data is empty")
	}
	return sll.Top.Data
}
