package testing

import (
	qll "data-structure/QueueLinkedList"
	"fmt"
	"testing"
)

func TestQueueLinkedList(t *testing.T) {
	queue := qll.Init()
	fmt.Println("Add Data: 5,6,4,1,9")
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(4)
	queue.Enqueue(1)
	queue.Enqueue(9)
	queue.DisplayData() // <- 5 <- 6 <- 4 <- 1 <- 9
	fmt.Println("Remove head/front: ", queue.Dequeue())
	fmt.Println("Remove head/front: ", queue.Dequeue())
	queue.DisplayData()                           // <- 4 <- 1 <- 9
	fmt.Println("Get head/front: ", queue.Peek()) // 4

}
