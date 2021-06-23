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

// fast, but ask for a lot of memory
// 06/23/2021 00:05	Accepted	4 ms	3.9 MB	golang
func twoSum(nums []int, target int) []int {
	// 进行了一次copy，需要额外的空间冗余
	newNums := make([]int, len(nums))
	copy(newNums, nums)
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
	testData := []struct {
		Input  []int
		Target int
		Dest   []int
	}{
		{
			Input:  getOrigin(),
			Target: 23,
			Dest:   []int{0, 4},
		},
		{
			Input:  []int{3, 3},
			Target: 6,
			Dest:   []int{0, 1},
		},
	}
	for _, datum := range testData {
		ret := twoSum(datum.Input, datum.Target)
		if !reflect.DeepEqual(datum.Dest, ret) {
			t.Error("err", datum.Input, datum.Target, ret)
		} else {
			t.Log("success", datum.Input, datum.Target, ret)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	testData := []struct {
		Input  []int
		Target int
		Dest   []int
	}{
		{
			Input:  getOrigin(),
			Target: 23,
			Dest:   []int{0, 4},
		},
		{
			Input:  []int{3, 3},
			Target: 6,
			Dest:   []int{0, 1},
		},
		{
			Input:  []int{3, 2, 4},
			Target: 6,
			Dest:   []int{1, 2},
		},
	}
	for _, datum := range testData {
		ret := twoSum2(datum.Input, datum.Target)
		if !reflect.DeepEqual(datum.Dest, ret) {
			t.Error("err", datum.Input, datum.Target, ret)
		} else {
			t.Log("success", datum.Input, datum.Target, ret)
		}
	}
}

// 06/23/2021 10:57	Accepted	4 ms	6.4 MB	golang
func twoSum2(nums []int, target int) []int {
	valIdxMap := make(map[int][]int)
	for idx, num := range nums {
		idxList, ok := valIdxMap[num]
		if !ok {
			valIdxMap[num] = []int{idx}
		} else {
			valIdxMap[num] = append(idxList, idx)
		}
	}
	for idx, num := range nums {
		idxList, ok := valIdxMap[target-num]
		if ok {
			if len(idxList) == 1 && num != target-num {
				return []int{idx, idxList[0]}
			} else if len(idxList) == 2 && num == target-num {
				return []int{idx, idxList[1]}
			}
		}
	}
	return nil
}
