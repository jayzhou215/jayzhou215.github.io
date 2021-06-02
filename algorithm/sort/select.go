package main

import (
	"fmt"
	"reflect"
)

var (
	unsortedArray = []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	sortedArray   = []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}
)

func main() {
	// select sort
	// 1. loop n-1 time
	// 2. pick min val idx, change swap the first val and min val
	selectSort()
	judge(unsortedArray)
}

func selectSort() {
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

func judge(array []int) {
	fmt.Println(reflect.DeepEqual(array, sortedArray), array)
}
