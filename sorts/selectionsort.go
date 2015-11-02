package sorts

//选择排序
func SelectionSort(data []int) {
	count := len(data)
	for i := 0; i < count-1; i++ {
		//假想最小值索引
		index := i
		for j := i + 1; j < count; j++ {
			//如果index下标的值大于j下标的值，则记录较小值下标j
			if data[index] > data[j] {
				index = j
			}
		}
		//最后将假想最小值跟最小的值进行交换
		data[index], data[i] = data[i], data[index]
	}
}
