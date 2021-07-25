package testing

import (
	bs "data-structure/BubbleSort"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{3, 7, 2, 9, 4, 5}

	fmt.Println(bs.BubbleSort(data))
}
