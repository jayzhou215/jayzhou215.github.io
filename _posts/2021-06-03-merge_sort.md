---
layout: post
title: Algorithm - merge sort
tags: [go]
readtime: true
comments: true
---

### merge sort
1. 合并排序的定义 是创建在归并操作上的一种有效的排序算法，效率为 O(n\log n)，该算法是采用分治法（Divide and Conquer）的一个非常典型的应用，且各层分治递归可以同时进行
2. 基本步骤
    1. 挑选基准值：从数列中挑出一个元素，称为“基准”（pivot），
    2. 分割：将array分割成ceil(n/2)，再切割成ceil(n/4)，直到包含1或2个元素
    3. 递归排序子序列：递归的排序再合并
2. 示例代码分析
    1. 分割成左右两部分，并且小于基准的在左，大于基准的在右
    2. 递归左右
3. [merge sort code](../algorithm/sort/sort_test.go)
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
BenchmarkBubble-12    	1000000000	         0.00104 ns/op
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.000895 ns/op
BenchmarkInsert
BenchmarkInsert-12    	1000000000	         0.000286 ns/op
BenchmarkQuick
BenchmarkQuick-12     	1000000000	         0.000093 ns/op
BenchmarkMerge
BenchmarkMerge-12     	1000000000	         0.000254 ns/op

# 10000个int数排序

BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.104 ns/op
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.0504 ns/op
BenchmarkInsert
BenchmarkInsert-12    	1000000000	         0.0188 ns/op
BenchmarkQuick
BenchmarkQuick-12     	1000000000	         0.000766 ns/op
BenchmarkMerge
BenchmarkMerge-12     	1000000000	         0.00114 ns/op


# 100000
BenchmarkBubble
BenchmarkBubble-12    	       1	13040344011 ns/op
BenchmarkSelect
BenchmarkSelect-12    	       1	5272985270 ns/op
BenchmarkInsert
BenchmarkInsert-12    	       1	1887316288 ns/op
BenchmarkQuick
BenchmarkQuick-12     	1000000000	         0.00894 ns/op
BenchmarkMerge
BenchmarkMerge-12     	1000000000	         0.0139 ns/op
```