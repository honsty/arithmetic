package sorts

import (
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	t.Skip()
	data := []int{50, 10, 30, 20, 40}
	t.Logf("排序前%#v", data)
	BubbleSort(data)
	t.Logf("冒泡排序后%#v", data)
}

func TestQuickSort(t *testing.T) {
	t.Skip()
	data := []int{20, 40, 50, 10, 60}
	t.Logf("排序前：%#v", data)
	QuickSort(data, 0, 4)
	t.Logf("快速排序后%#v", data)
}

func TestSelectSort(t *testing.T) {
	t.Skip()
	data := []int{20, 40, 50, 10, 60}
	t.Logf("排序前：%#v", data)
	SelectionSort(data)
	t.Logf("选择排序后%#v", data)
}

func TestHeapSort(t *testing.T) {
	for j := 1; j < 5; j++ {
		list := make([]int, 0, 20000)
		for i := 0; i < 20000; i++ {
			list = append(list, rand.Intn(99999))
		}
		now := time.Now()
		QuickSort(list, 0, 19999)
		t.Logf("第 %d 次快速排序时间是:%d,%#v", j, time.Now().UnixNano()-now.UnixNano(), list[:10])

		now = time.Now()
		HeapSort(list)
		t.Logf("第 %d 次堆排序时间是:%d,%#v", j, time.Now().UnixNano()-now.UnixNano(), list[:10])
	}
}

func TestHeapSortTake(t *testing.T) {
	for j := 1; j < 5; j++ {
		list := make([]int, 0, 20000)
		for i := 0; i < 20000; i++ {
			list = append(list, rand.Intn(90000))
		}
		now := time.Now()
		QuickSort(list, 0, 19999)
		t.Logf("第 %d 次快速排序前10时间是:%d,%#v", j, time.Now().UnixNano()-now.UnixNano(), list[:10])

		now = time.Now()
		data := HeapSortTake(list, 10)
		t.Logf("第 %d 次堆排序前10时间是:%d,%#v", j, time.Now().UnixNano()-now.UnixNano(), data[:10])
	}
}

func TestInsertSort(t *testing.T) {
	data := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, rand.Intn(1000000))
	}
	t.Logf("直接插入排序前:%#v\n", data[:10])
	InsertSort(data)
	t.Logf("直接插入排序后:%#v\n", data[:10])
}

func TestShellSort(t *testing.T) {
	data := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, rand.Intn(1000000))
	}
	t.Logf("希尔排序前:%#v\n", data[:10])
	ShellSort(data)
	t.Logf("希尔排序后:%#v\n", data[:10])
}

func TestMerSort(t *testing.T) {
	arr := []int{3, 2, 1, 8, 9, 0}
	temp := make([]int, len(arr))
	MergeSort(arr, temp, 0, len(arr)-1)
	t.Logf("arr=%#v", arr)
}
