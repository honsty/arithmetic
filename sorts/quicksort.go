package sorts

//分割函数
func division(data []int, left, right int) int {
	base := data[left]
	for left < right {
		for left < right && data[right] >= base {
			right = right - 1
		}

		data[left] = data[right]

		for left < right && data[left] < base {
			left += 1
		}

		data[right] = data[left]
	}
	data[left] = base
	return left
}

//快速排序
func QuickSort(data []int, left, right int) {
	if left < right {
		i := division(data, left, right)

		//对左侧排序
		QuickSort(data, left, i-1)

		//对右侧排序
		QuickSort(data, i+1, right)
	}
}
