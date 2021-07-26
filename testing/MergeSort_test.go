package testing

import (
	bs "data-structure/MergeSort"
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := []int{17, 12, 29, 3, 14, 76, 55}
	result := bs.MergeSort(data)
	fmt.Println(result)

	expected := []int{3, 12, 14, 17, 29, 55, 76}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\n got %v \n want %v \n given %v", result, expected, data)
	}
}
