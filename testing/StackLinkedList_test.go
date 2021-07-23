package testing

import (
	sll "data-structure/StackLinkedList"
	"fmt"
	"testing"
)

func TestStackLinkedList(t *testing.T) {
	stack := sll.Init()
	fmt.Println("Push 9 -> 2 -> 7 -> 4")
	stack.Push(4)                          // 4
	stack.Push(7)                          // 7 -> 4
	stack.Push(2)                          // 2 -> 7 -> 4
	stack.Push(9)                          // 9 -> 2 -> 7 -> 4
	fmt.Println("Size: ", stack.GetSize()) // 4
	stack.DisplayData()
	fmt.Println("Pop: ", stack.Pop()) // 2 -> 7 -> 4
	stack.DisplayData()
	fmt.Println("Pop: ", stack.Pop()) // 7 -> 4
	fmt.Println("Push 8")
	stack.Push(8)                          // 8 -> 7 -> 4
	fmt.Println("Peek: ", stack.Peek())    // 8
	stack.DisplayData()                    // 8 -> 7 -> 4
	fmt.Println("Pop: ", stack.Pop())      // 8
	fmt.Println("Size: ", stack.GetSize()) // 2
	fmt.Println("Pop: ", stack.Pop())      // 7
	fmt.Println("Pop: ", stack.Pop())      // 4
	fmt.Println("Pop: ", stack.Pop())      // Data is Empty
	stack.DisplayData()                    // Data is Empty
}
