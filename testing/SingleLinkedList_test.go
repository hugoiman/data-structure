package main

import (
	sll "data-structure/SingleLinkedList"
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	list := sll.Init()
	fmt.Printf("Add Front: %d, --> ", 4)
	list.AddFirst(4)
	list.DisplayData() // 4

	fmt.Printf("Add Front: %d, --> ", 2)
	list.AddFirst(2)
	list.DisplayData() // 2, 4

	fmt.Printf("Add Front: %d, --> ", 6)
	list.AddFirst(6)
	list.DisplayData() // 6, 2, 4

	fmt.Printf("Add Last: %d, --> ", 9)
	list.AddLast(9)
	list.DisplayData() // 6, 2, 4, 9

	fmt.Printf("Add Last: %d, --> ", 19)
	list.AddLast(19)
	list.DisplayData() // 6, 2, 4, 9, 19

	fmt.Printf("Add %d After %d, --> ", 5, 4)
	list.InsertAfter(5, 4)
	list.DisplayData() // 6, 2, 4, 5, 9, 19

	fmt.Printf("Add %d After %d, --> ", 13, 11)
	list.InsertAfter(13, 11) // Data 11 not found
	list.DisplayData()       // 6, 2, 4, 5, 9, 19

	fmt.Printf("Remove Front --> ")
	list.RemoveFirst()
	list.DisplayData() // 2, 4, 5, 9, 19

	fmt.Printf("Remove Last --> ")
	list.RemoveLast()
	list.DisplayData() // 2, 4, 5, 9

	fmt.Printf("Remove: %d, --> ", 4)
	list.RemoveCertain(4)
	list.DisplayData() // 2, 5, 9

	fmt.Printf("Remove: %d, --> ", 12)
	list.RemoveCertain(12) // Data 12 not found
	list.DisplayData()     // 2, 5, 9

	fmt.Printf("Get data at index %d: ", 1)
	data, err := list.GetIndex(1)
	fmt.Println(data, err) // 5

	fmt.Printf("Get data at index %d: ", 4)
	data, err = list.GetIndex(4)
	fmt.Println(data, err) // index 4 out of bounds exception
}
