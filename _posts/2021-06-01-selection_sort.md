---
layout: post
title: Algorithm - select sort
tags: [go]
readtime: true
comments: true
---

### select sort
1. 选择排序的定义 它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢"浮"到数列的顶端。
2. 基本步骤
    1. 先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置　
    2. 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
    3. 以此类推，直到所有元素均排序完毕
2. 示例代码分析
    1. 循环n-1次，此循环次数不可减少
    2. 记录单次循环中最小值位置，交互未排序数组中首元素值与最小值
3. [select sort code](../algorithm/sort/select.go)
4. 时间空间复杂度
    1. 时间复杂度: 外循环和内循环以及判断和交换元素的时间开销
        * 比较操作，总是 n*(n-1)/2，即O(n^2)
        * 交换操作，最优0次，最坏n-1次，即O(n)
        * 赋值操作，0 ~ 3*(n-1)
    2. 空间复杂度: 交换元素时那个临时变量所占的内存空间
        * 交互在golang中是隐式的，空间为2，即O(1)
        * 赋值操作空间为1，即O(1)
    3. 稳定性 stability
        * 因为判断条件是left > right，等值时不会发生交换，总是稳定的    

### 优点
简单 & in-place(原地操作，无需过多辅助空间)
比冒泡排序快的原因
### 引申知识点
1. 交换次数比冒泡排序较少，由于交换所需CPU时间比比较所需的CPU时间多，n值较小时，选择排序比冒泡排序快。
2. 当n值较大时，选择排序明显慢于冒泡排序，原因是冒泡排序通过判断是否有发生交换来跳出循环，但选择排序无法做出类似的优化

### bench测试

```go
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

```

```sh
# 10000个int数排序
goos: darwin
goarch: amd64
BenchmarkSelect
BenchmarkSelect-12    	       1	5187351688 ns/op
BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.000102 ns/op

# 1000个int数排序
goos: darwin
goarch: amd64
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.000731 ns/op
BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.000002 ns/op
```