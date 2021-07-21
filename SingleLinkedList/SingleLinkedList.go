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

func (ll *SingleLinkedList) IsEmpty() bool {
	return (ll.Head == SingleLinkedList{}.Head)
}

func (ll *SingleLinkedList) GetSize() int {
	return ll.Size
}

func (ll *SingleLinkedList) DisplayData() {
	temp := ll.Head
	fmt.Print("Data List: ")
	for temp != nil {
		fmt.Printf("%d, ", temp.Data)
		temp = temp.Next
	}
	fmt.Println(fmt.Sprintf("Head: %d, Tail: %d, Size: %d ", ll.Head.Data, ll.Tail.Data, ll.Size))
}

func (ll *SingleLinkedList) AddFirst(input int) {
	newNode := new(Node)
	newNode.Node(input)
	if ll.IsEmpty() {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = newNode
	}
	ll.Size++
}

func (ll *SingleLinkedList) AddLast(input int) {
	newNode := new(Node)
	newNode.Node(input)
	if ll.IsEmpty() {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++
}

func (ll *SingleLinkedList) InsertAfter(input, data int) {
	newNode := new(Node)
	newNode.Node(input)
	temp := ll.Head

	for temp != nil {
		if temp.Data == data {
			newNode.Next = temp.Next
			temp.Next = newNode
			ll.Size++
			break
		} else if temp == ll.Tail && temp.Data != data {
			fmt.Printf("Data %d not found. ", data)
		}
		temp = temp.Next
	}
}

func (ll *SingleLinkedList) RemoveFirst() {
	temp := ll.Head
	if !ll.IsEmpty() {
		if ll.Head == ll.Tail {
			ll.Head, ll.Tail = nil, nil
		} else {
			ll.Head = temp.Next
		}
		ll.Size--
	} else {
		fmt.Print("Data is empty")
	}
}

func (ll *SingleLinkedList) RemoveLast() {
	temp := ll.Head
	if !ll.IsEmpty() {
		if ll.Tail == ll.Head {
			ll.Head, ll.Tail = nil, nil
		} else {
			for temp.Next != ll.Tail {
				temp = temp.Next
			}
			temp.Next = nil
			ll.Tail = temp
		}
		ll.Size--
	} else {
		fmt.Print("Data is empty")
	}
}

func (ll *SingleLinkedList) RemoveCertain(data int) {
	temp := ll.Head
	if !ll.IsEmpty() {
		for temp != nil {
			if temp == ll.Tail && temp.Data != data {
				fmt.Printf("Data %d not found. ", data)
				break
			} else if temp.Next.Data == data {
				temp.Next = temp.Next.Next
				if temp.Next == nil {
					ll.Tail = temp
				}
				ll.Size--
				break
			} else if temp.Data == data && temp == ll.Head {
				ll.RemoveFirst()
				ll.Size--
				break
			}
			temp = temp.Next
		}
	} else {
		fmt.Print("Data is empty")
	}
}

func (ll *SingleLinkedList) CheckIndex(index int) {
	if index < 0 || index > ll.Size {
		fmt.Printf("index %d out of bounds exception ", index)
		os.Exit(0)
	}
}

func (ll *SingleLinkedList) GetIndex(index int) int {
	ll.CheckIndex(index)
	temp := ll.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp.Data
}
