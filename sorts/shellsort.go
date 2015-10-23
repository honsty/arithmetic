package sorts

//希尔排序
func ShellSort(data []int) {
	count := len(data)
	step := count / 2
	for step >= 1 {
		for i := step; i < count; i++ {
			temp := data[i]
			var j int
			for j = i - step; j >= 0 && temp < data[j]; j = j - step {
				data[j+step] = data[j]
			}
			data[j+step] = temp
		}
		step = step / 2
	}
}
