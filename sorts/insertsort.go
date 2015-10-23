package sorts

//插入排序
func InsertSort(data []int) {
	count := len(data)
	for i := 1; i < count; i++ {
		temp := data[i]
		var j int
		for j = i - 1; j >= 0 && temp < data[j]; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}
}
