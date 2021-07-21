package main

import (
	"fmt"

	sll "data-structure/SingleLinkedList"
)

func main() {
	list := sll.Init()

	fmt.Printf("Add Front: %d, --> ", 4)
	list.AddFirst(4)
	list.DisplayData()

	fmt.Printf("Add Front: %d, --> ", 2)
	list.AddFirst(2)
	list.DisplayData()

	fmt.Printf("Add Front: %d, --> ", 6)
	list.AddFirst(6)
	list.DisplayData()

	fmt.Printf("Add Last: %d, --> ", 9)
	list.AddLast(9)
	list.DisplayData()

	fmt.Printf("Add Last: %d, --> ", 19)
	list.AddLast(19)
	list.DisplayData()

	fmt.Printf("Add %d After %d, --> ", 5, 4)
	list.InsertAfter(5, 4)
	list.DisplayData()

	fmt.Printf("Add %d After %d, --> ", 13, 11)
	list.InsertAfter(13, 11)
	list.DisplayData()

	fmt.Printf("Remove Front --> ")
	list.RemoveFirst()
	list.DisplayData()

	fmt.Printf("Remove Last --> ")
	list.RemoveLast()
	list.DisplayData()

	fmt.Printf("Remove: %d, --> ", 4)
	list.RemoveCertain(4)
	list.DisplayData()

	fmt.Printf("Remove: %d, --> ", 12)
	list.RemoveCertain(12)
	list.DisplayData()

	fmt.Printf("Get data at index %d: ", 1)
	fmt.Println(list.GetIndex(1))
	fmt.Printf("Get data at index %d: ", 4)
	fmt.Println(list.GetIndex(4))
}
