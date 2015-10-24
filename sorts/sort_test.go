package sorts

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{50, 10, 30, 20, 40}
	t.Logf("排序前%#v", data)
	BubbleSort(data)
	t.Logf("冒泡排序后%#v", data)
}

func TestQuickSort(t *testing.T) {
	data := []int{20, 40, 50, 10, 60}
	t.Logf("排序前：%#v", data)
	QuickSort(data, 0, 4)
	t.Logf("快速排序后%#v", data)
}

func TestSelectSort(t *testing.T) {
	data := []int{20, 40, 50, 10, 60}
	t.Logf("排序前：%#v", data)
	SelectionSort(data)
	t.Logf("选择排序后%#v", data)
}
