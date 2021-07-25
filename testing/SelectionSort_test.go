package testing

import (
	bs "data-structure/SelectionSort"
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := []int{3, 7, 2, 9, 4, 5}

	fmt.Println(bs.SelectionSort(data))
}
