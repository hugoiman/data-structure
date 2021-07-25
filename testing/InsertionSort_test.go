package testing

import (
	bs "data-structure/InsertionSort"
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	data := []int{3, 7, 2, 9, 4, 5}

	fmt.Println(bs.InsertionSort(data))
}
