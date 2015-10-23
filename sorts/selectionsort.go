package sorts

//选择排序
func SelectionSort(data []int) {
	count := len(data)
	for i := 0; i < count-1; i++ {
		index := i
		for j := i + 1; j < count; j++ {
			if data[index] > data[j] {
				index = j
			}
		}
		data[index], data[i] = data[i], data[index]
	}
}
