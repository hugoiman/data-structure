package testing

import (
	bs "data-structure/SelectionSort"
	"fmt"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := []int{3, 7, 2, 9, 4, 5}
	result := bs.SelectionSort(data)
	fmt.Println(result)

	expected := []int{2, 3, 4, 5, 7, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\n got %v \n want %v \n given %v", result, expected, data)
	}
}
