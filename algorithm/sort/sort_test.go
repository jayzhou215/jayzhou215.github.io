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
	for i := 0; i < 10000; i++ {
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

func BenchmarkSelect(b *testing.B) {
	newArray := initArray()
	selectSort(newArray)
}

func BenchmarkBubble(b *testing.B) {
	newArray := initArray()
	bubbleSort(newArray)
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
