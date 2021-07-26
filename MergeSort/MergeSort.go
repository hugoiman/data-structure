package Mergesort

func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	} else {
		middle := len(data) / 2
		left := MergeSort(data[:middle])
		right := MergeSort(data[middle:])
		return Merge(left, right)
	}
}

func Merge(left []int, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
