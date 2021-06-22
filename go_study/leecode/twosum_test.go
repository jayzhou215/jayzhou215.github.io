package leecode

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	sorted = []int{1, 3, 7, 9, 18, 22, 33, 34, 55}
)

func getOrigin() []int {
	return []int{1, 9, 3, 7, 22, 18, 34, 33, 55}
}

func twoSum(nums []int, target int) []int {
	// 进行了一次copy，需要额外的空间冗余
	newNums := make([]int, 0)
	newNums = append(newNums, nums...)
	// sort
	quickSort(newNums, 0, len(newNums)-1)
	// pick出对应的值
	start := 0
	end := len(newNums) - 1
	for start < end {
		if newNums[start]+newNums[end] > target {
			end--
		} else if newNums[start]+newNums[end] < target {
			start++
		} else {
			break
		}
	}
	startVal := newNums[start]
	endVal := newNums[end]
	fmt.Println(newNums, start, end, target)
	// 再次遍历，找到索引，需考虑 相同值的情况，左右逼近原则
	start = 0
	end = len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if startVal == nums[i] {
			start = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if endVal == nums[i] {
			end = i
			break
		}
	}
	return []int{start, end}
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	// 随机取一个锚点
	pivot := (start + end) / 2
	// 将所有比锚点值大的值放到前面，定位到新的索引位置，将新的索引位置记录为新的锚点位置
	newPivot := getNewPivot(nums, start, pivot, end)
	quickSort(nums, start, newPivot-1)
	quickSort(nums, newPivot+1, end)
}

func getNewPivot(nums []int, start int, pivot int, end int) int {
	pivotVal := nums[pivot]
	swap(nums, pivot, end)
	newPivot := start
	for i := start; i < end; i++ {
		if nums[i] < pivotVal {
			swap(nums, i, newPivot)
			newPivot++
		}
	}
	swap(nums, end, newPivot)
	return newPivot
}

func swap(nums []int, idxA, idxB int) {
	nums[idxA], nums[idxB] = nums[idxB], nums[idxA]
}

func TestQuickSort(t *testing.T) {
	origin := getOrigin()
	quickSort(origin, 0, len(origin)-1)
	if reflect.DeepEqual(origin, sorted) {
		t.Log("ok", origin, sorted)
	} else {
		t.Error("fail", origin, sorted)
	}
}

func TestTwoSum(t *testing.T) {
	origin := getOrigin()
	target := 23
	hope := []int{0, 4}
	ret := twoSum(origin, target)
	if !reflect.DeepEqual(hope, ret) {
		t.Error("err", origin, target, ret)
	} else {
		t.Log("success", origin, target, ret)
	}
}
