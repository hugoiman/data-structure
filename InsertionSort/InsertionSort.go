package InsertionSort

func InsertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				temp := data[j]
				data[j] = data[j-1]
				data[j-1] = temp
			}
		}
	}
	return data
}
