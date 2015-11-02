package sorts

//分割函数
func division(data []int, left, right int) int {
	//基准元素
	base := data[left]
	for left < right {
		//从右端向前找，一直找到比Base小的为止
		for left < right && data[right] >= base {
			right = right - 1
		}

		//找到比Base小的元素，将此元素放到base的位置
		data[left] = data[right]

		//从左端向后找，一直找到比Base大的为止
		for left < right && data[left] < base {
			left += 1
		}
		//找到了比Base大的元素，将此元素放到最后的位置
		data[right] = data[left]
	}
	//最后将Base放到left位置
	data[left] = base

	//发现，left位置左侧部分都比left小,右侧部分都比left大
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
