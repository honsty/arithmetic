package sorts

/**构建堆
@param data []int  待排序
@param parent int  父节点
@param length int  输出根堆时剔除最大值使用
*/
func heapAdjust(data []int, parent, length int) {
	//保存当前父节点
	temp := data[parent]
	//得到左孩子，
	child := 2*parent + 1

	for child < length {
		//如果parent有右孩子，则要判断左孩子是否小于右孩子
		if child+1 < length && data[child] < data[child+1] {
			child++
		}
		//你节点大于子节点，就不用做交换
		if temp >= data[child] {
			break
		}

		//将较大节点的值赋给父节点
		data[parent] = data[child]

		parent = child

		child = 2*parent + 1
	}
	data[parent] = temp
}

//堆排序
func HeapSort(data []int) {
	count := len(data)
	for i := count/2 - 1; i >= 0; i-- {
		heapAdjust(data, i, count)
	}

	//最后输入堆元素
	for i := count - 1; i > 0; i-- {

		//堆顶与当前堆的第i个元素进行值对调
		data[0] = data[i]

		heapAdjust(data, 0, i)
	}

}
