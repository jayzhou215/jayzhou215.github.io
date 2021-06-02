package main

import (
	"math/rand"
	"testing"
)

var (
	unsortedArray []int
)

func init() {
	for i := 0; i < 1000; i++ {
		randInt := rand.Int()
		unsortedArray = append(unsortedArray, randInt)
	}
}

func BenchmarkSelect(b *testing.B) {
	selectSort()
}

func BenchmarkBubble(b *testing.B) {
	bubbleSort()
}

func bubbleSort() {
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
