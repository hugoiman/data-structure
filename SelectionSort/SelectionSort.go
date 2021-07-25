package SelectionSort

func SelectionSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				temp := data[j]
				data[j] = data[i]
				data[i] = temp
			}
		}
	}
	return data
}
