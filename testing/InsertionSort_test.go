package testing

import (
	bs "data-structure/InsertionSort"
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	data := []int{3, 7, 2, 9, 4, 5}
	result := bs.InsertionSort(data)
	fmt.Println(result)

	expected := []int{2, 3, 4, 5, 7, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\n got %v \n want %v \n given %v", result, expected, data)
	}
}
