package sorts

//冒泡排序
func BubbleSort(data []int) {
	count := len(data)

	for i := 0; i < count-1; i++ {
		//从后往前的下标一定大于从前往后的下标
		for j := count - 1; j > i; j-- {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}
}
