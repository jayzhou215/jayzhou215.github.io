package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

var (
	sortedArray = []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}
)

func initArray() []int {
	var unsortedArray []int
	for i := 0; i < 1000; i++ {
		randInt := rand.Int()
		unsortedArray = append(unsortedArray, randInt)
	}
	return unsortedArray
}

func TestBubble(t *testing.T) {
	unsortedArray := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	bubbleSort(unsortedArray)
	judge(unsortedArray, sortedArray)

}

func TestSelect(t *testing.T) {
	unsortedArray := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	selectSort(unsortedArray)
	judge(unsortedArray, sortedArray)

}

func TestInsert(t *testing.T) {
	unsortedArray := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	insertSort(unsortedArray)
	judge(unsortedArray, sortedArray)
}

func TestQuick(t *testing.T) {
	unsortedArray := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	quickSort(unsortedArray, 0, len(unsortedArray)-1)
	judge(unsortedArray, sortedArray)
}

func quickSort(unsortedArray []int, left, right int) {
	if right > left {
		// 随机选择一个基准点，这里取中间的值，可以随意
		pivotIdx := (right + left) / 2
		// 分割array, 所有小于基准值的在左侧，大于的在右侧
		pivotNewIdx := partition(unsortedArray, left, right, pivotIdx)
		quickSort(unsortedArray, left, pivotNewIdx-1)
		quickSort(unsortedArray, pivotNewIdx+1, right)
	}
}

func partition(array []int, left int, right int, pivotIdx int) (storeIdx int) {
	pivotVal := array[pivotIdx]
	// 从左一开始记做存储位置
	storeIdx = left
	// 先将pivot缓存到right位置
	swap(array, pivotIdx, right)
	// 从left开始，一直到right-1，包括right-1
	for i := left; i <= right-1; i++ {
		// 所有小的依次放到左一左二...
		if array[i] < pivotVal {
			swap(array, i, storeIdx)
			storeIdx++
		}
	}
	// 依次放置完，最终处理下之前缓存过的pivot
	swap(array, right, storeIdx)
	return
}

func insertSort(unsortedArray []int) {
	// 第一位默认有序
	// 每次将新的元素插入到有序序列中
	for i := 1; i < len(unsortedArray); i++ {
		// 从无序数组中取出首位，待插入有序数组
		current := unsortedArray[i]

		// 记录当前待插入索引，默认 -1
		insertIdx := -1

		// 从0到i-1为有序数组，倒序遍历
		for j := i - 1; j >= 0; j-- {
			// 有序数组的第j个值如比current大，则将第j位值后移一位
			if unsortedArray[j] > current {
				// 记录待插入索引
				insertIdx = j
				unsortedArray[j+1] = unsortedArray[j]
			} else {
				// 找到正确位置，跳出当前循环
				break
			}
		}

		// 有效的索引值 > -1，找到需要更新的位置，进行插入
		if insertIdx > -1 {
			unsortedArray[insertIdx] = current
		}

	}
}

func BenchmarkBubble(b *testing.B) {
	newArray := initArray()
	bubbleSort(newArray)
}

func BenchmarkSelect(b *testing.B) {
	newArray := initArray()
	selectSort(newArray)
}

func BenchmarkInsert(b *testing.B) {
	newArray := initArray()
	insertSort(newArray)
}

func BenchmarkQuick(b *testing.B) {
	newArray := initArray()
	quickSort(newArray, 0, len(newArray)-1)
}

func bubbleSort(unsortedArray []int) {
	swapped := true
	lastUnsortedElementIndex := len(unsortedArray)
	for swapped {
		swapped = false
		for i := 0; i < lastUnsortedElementIndex-1; i++ {
			if unsortedArray[i] > unsortedArray[i+1] {
				swap(unsortedArray, i, i+1)
				swapped = true
			}
		}
		lastUnsortedElementIndex--
		if !swapped {
			break
		}
	}
}

func selectSort(unsortedArray []int) {
	for i := 0; i < len(unsortedArray)-1; i++ {
		minValIdx := i
		for j := i + 1; j < len(unsortedArray); j++ {
			if unsortedArray[minValIdx] > unsortedArray[j] {
				minValIdx = j
			}
		}
		if minValIdx != i {
			swap(unsortedArray, i, minValIdx)
		}
	}
}

func swap(array []int, leftIndex int, rightIndex int) {
	array[leftIndex], array[rightIndex] = array[rightIndex], array[leftIndex]
}

func judge(array []int, sortedArray []int) {
	fmt.Println(reflect.DeepEqual(array, sortedArray), array)
}
