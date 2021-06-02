package main

import (
	"fmt"
	"reflect"
)

func main() {
	unsortedArray := []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}
	// 首位默认为已排序
	// 从idx=1开始，处理无序数组
	for i := 1; i < len(unsortedArray); i++ {
		// 新增记录插入位置idx变量
		insertIdx := -1

		// 新增当前无序数组首位值变量
		firstUnsortVal := unsortedArray[i]

		// 倒序遍历有序数组
		for j := i - 1; j >= 0; j-- {
			// 有序数组的当前值 如 > 无序数组首位值，则将当前值右移一位，直至ok
			if unsortedArray[j] > firstUnsortVal {
				unsortedArray[j+1] = unsortedArray[j]
				insertIdx = j
			} else {
				break
			}
		}

		// 跳出循环后，如果insertIdx 不等于默认值，则将该位置的值设为无序数组首位值
		if insertIdx > -1 {
			unsortedArray[insertIdx] = firstUnsortVal
		}
	}

	sortedArray := []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}
	judge(unsortedArray, sortedArray)
}

func judge(array []int, sortedArray []int) {
	fmt.Println(reflect.DeepEqual(array, sortedArray), array)
}
