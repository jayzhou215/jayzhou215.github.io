---
layout: post
title: Algorithm - quick sort
tags: [go]
readtime: true
comments: true
---

### quick sort
1. 快速排序的定义 使用分治法（Divide and conquer，直译是分割击破）策略来把一个序列（list）分为较小和较大的2个子序列，然后递归地排序两个子序列。
2. 基本步骤
    1. 挑选基准值：从数列中挑出一个元素，称为“基准”（pivot），
    2. 分割：重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（与基准值相等的数可以到任何一边）。在这个分割结束之后，对基准值的排序就已经完成，
    3. 递归排序子序列：递归地将小于基准值元素的子序列和大于基准值元素的子序列排序。
2. 示例代码分析
    1. 分割成左右两部分，并且小于基准的在左，大于基准的在右
    2. 递归左右
3. [quick sort code](../algorithm/sort/sort_test.go)
4. 时间空间复杂度
    1. 时间复杂度:
        * 最优O(n*log(n))，最坏O(n^2)
    2. 空间复杂度: 
        * 交互在golang中是隐式的，空间为2，即O(1)
    3. 稳定性 stability
        * unstable    

### 优点
最优O(n*log(n))的速度

### 缺点
unstable

### 引申知识点
1. 递归的使用
2. in-place的使用节省了空间

### bench测试
[code](../algorithm/sort/sort_test.go)

```sh

# 1000个int数排序
BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.000865 ns/op
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.000779 ns/op
BenchmarkInsert
BenchmarkInsert-12    	1000000000	         0.000301 ns/op
BenchmarkQuick
BenchmarkQuick-12     	1000000000	         0.000138 ns/op

# 10000个int数排序

BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.102 ns/op
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.0511 ns/op
BenchmarkInsert
BenchmarkInsert-12    	1000000000	         0.0215 ns/op
BenchmarkQuick
BenchmarkQuick-12     	1000000000	         0.000843 ns/op

```