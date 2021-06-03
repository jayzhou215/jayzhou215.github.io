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

const (
	length = 1000000
)

func initArray() []int {
	var unsortedArray []int
	for i := 0; i < length; i++ {
		randInt := rand.Int()
		//randInt := rand.Intn(10000)
		unsortedArray = append(unsortedArray, randInt)
	}
	return unsortedArray
}

func initRandNArray() []int {
	var unsortedArray []int
	for i := 0; i < length; i++ {
		randInt := rand.Intn(length)
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

func TestMerge(t *testing.T) {
	unsortedArray := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	lastArray := mergeSort(unsortedArray)
	judge(lastArray, sortedArray)
}

func TestCount(t *testing.T) {
	unsortedArray := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	countSort(unsortedArray)
	judge(unsortedArray, sortedArray)
}

func TestRadix(t *testing.T) {
	unsortedRadixArray := []int{2435, -4224, 138, 5, 132147, 15, 346, -7, 27, 2, 2146, 12314, 19, 50, 48}
	sortedRadixArray := []int{-4224, -7, 2, 5, 15, 19, 27, 48, 50, 138, 346, 2146, 2435, 12314, 132147}
	radixSort(unsortedRadixArray)
	judge(unsortedRadixArray, sortedRadixArray)
}

func radixSort(array []int) {
	hasHigher := true
	level := 10
	for hasHigher {
		// to support negative [-9 ~ 9] + 9 => [0 ~ 18]
		buckets := make([][]int, 19)
		hasHigher = false
		// put array value into buckets
		for i := 0; i < len(array); i++ {
			val := array[i]
			digit := val%level/(level/10) + 9
			if val/level > 1 {
				hasHigher = true
			}
			list := buckets[digit]
			if list != nil {
				buckets[digit] = append(buckets[digit], val)
			} else {
				buckets[digit] = []int{val}
			}
		}
		// put back value into array
		idx := 0
		for _, list := range buckets {
			for _, val := range list {
				array[idx] = val
				idx++
			}
		}
		// inc level
		level *= 10
	}
}

func countSort(array []int) {
	// get max value in array
	maxVal := array[0]
	for i := 0; i < len(array); i++ {
		if maxVal < array[i] {
			maxVal = array[i]
		}
	}
	// declare countArr with len maxVal+1
	countArr := make([]int, maxVal+1)
	// incr val count
	for i := 0; i < len(array); i++ {
		val := array[i]
		countArr[val] += 1
	}
	idx := 0
	for val := 0; val < len(countArr); val++ {
		valCnt := countArr[val]
		for j := 0; j < valCnt; j++ {
			array[idx] = val
			idx++
		}
	}
}

func TestInPlaceMerge(t *testing.T) {
	unsortedArray := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	inPlaceMergeSort(unsortedArray, 0, len(unsortedArray))
	judge(unsortedArray, sortedArray)
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	// 先拆
	leftArr := mergeSort(array[0 : len(array)/2])
	rightArr := mergeSort(array[len(array)/2:])
	// 再合
	return merge(leftArr, rightArr)
}

func merge(arr1 []int, arr2 []int) []int {
	newArr := make([]int, len(arr1)+len(arr2))
	idx := 0
	idx1 := 0
	idx2 := 0
	for {
		// arr1 到达末尾
		if idx1 >= len(arr1) {
			// arr2还有数据
			if idx2 < len(arr2) {
				copy(newArr[idx:], arr2[idx2:])
			}
			break
		}
		// arr2 到达末尾
		if idx2 >= len(arr2) {
			// arr1还有数据
			if idx1 < len(arr1) {
				copy(newArr[idx:], arr1[idx1:])
			}
			break
		}
		// 两者都还有数据
		if arr1[idx1] < arr2[idx2] {
			newArr[idx] = arr1[idx1]
			idx++
			idx1++
		} else {
			newArr[idx] = arr2[idx2]
			idx++
			idx2++
		}
	}
	return newArr
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

func randomQuickSort(unsortedArray []int, left, right int) {
	if right > left {
		// 随机选择一个基准点，这里取中间的值，可以随意
		pivotIdx := rand.Intn(right-left) + left
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
		firstUnsortedVal := unsortedArray[i]

		// 记录当前待插入索引，默认 -1
		insertIdx := i

		// 从0到i-1为有序数组，倒序遍历
		for j := i - 1; j >= 0; j-- {
			// 有序数组的第j个值如比firstUnsortedVal大，则将第j位值后移一位
			if unsortedArray[j] > firstUnsortedVal {
				// 记录待插入索引
				insertIdx = j
				unsortedArray[j+1] = unsortedArray[j]
			} else {
				// 找到正确位置，跳出当前循环
				break
			}
		}

		// 找到需要更新的位置，进行插入
		if insertIdx != i {
			unsortedArray[insertIdx] = firstUnsortedVal
		}

	}
}

//func BenchmarkBubble(b *testing.B) {
//	newArray := initArray()
//	bubbleSort(newArray)
//}
//
//func BenchmarkSelect(b *testing.B) {
//	newArray := initArray()
//	selectSort(newArray)
//}
//
//func BenchmarkInsert(b *testing.B) {
//	newArray := initArray()
//	insertSort(newArray)
//}

func BenchmarkQuick(b *testing.B) {
	newArray := initArray()
	quickSort(newArray, 0, len(newArray)-1)
}

func BenchmarkRandomQuick(b *testing.B) {
	newArray := initArray()
	randomQuickSort(newArray, 0, len(newArray)-1)
}

func BenchmarkMerge(b *testing.B) {
	newArray := initArray()
	_ = mergeSort(newArray)
}

func BenchmarkRadix(b *testing.B) {
	newArray := initArray()
	radixSort(newArray)
}

//func BenchmarkCount(b *testing.B) {
//	newArray := initRandNArray()
//	countSort(newArray)
//}

//
//func BenchmarkInPlaceMerge(b *testing.B) {
//	newArray := initArray()
//	inPlaceMergeSort(newArray, 0, len(newArray))
//}

func inPlaceMergeSort(array []int, left int, right int) {
	if left+1 >= right {
		return
	}
	mid := (left + right) / 2
	inPlaceMergeSort(array, left, mid)
	inPlaceMergeSort(array, mid, right)
	inPlaceMerge(array, left, mid, right)

}

func inPlaceMerge(array []int, left int, mid int, right int) {
	// merge arr1=array[left: mid), arr2=array[mid: right) in place
	// get min value from arr1, arr2, and then put it into sorted array
	// define sortedIdx, arr1Idx, arr2Idx
	arr1Idx := left
	arr2Idx := mid
	for arr1Idx < right && arr2Idx < right {
		// pick a small one to i
		if array[arr1Idx] < array[arr2Idx] {
			arr1Idx++
		} else {
			// insert arr2Idx value into insertIdx
			tmp := array[arr2Idx]
			insertIdx := arr2Idx
			for j := arr2Idx - 1; j >= arr1Idx; j-- {
				if array[j] > tmp {
					array[j+1] = array[j]
					insertIdx = j
				} else {
					break
				}
			}
			arr2Idx++
			array[insertIdx] = tmp
			arr1Idx++
		}
	}
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
	// iterate [0: n-1)
	// minIdx = i
	// iterate unsorted array [i+1, n)
	// minIdx = j if arr[minIdx] > arr[j]
	// after internal iterate swap(arr[i], arr[minIdx])
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
