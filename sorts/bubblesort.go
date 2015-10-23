package sorts

func BubbleSort(data []int) {
	count := len(data)
	for i := 0; i < count-1; i++ {
		for j := count - 1; j > i; j-- {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}
}
