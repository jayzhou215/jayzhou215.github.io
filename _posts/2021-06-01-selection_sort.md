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
简单 & in-place(原地操作，无需过多辅助空间) & 大大减少了交换次数

### 缺点
1、依然存在大量无效比较，只在当次循环

### 引申知识点
1. 由于交换所需CPU时间比比较所需的CPU时间多

### bench测试
[code](../algorithm/sort/sort_test.go)

```sh
# 10000个int数排序
BenchmarkSelect
BenchmarkSelect-12    	     381	   3337452 ns/op
BenchmarkBubble
BenchmarkBubble-12    	       1	2639470895 ns/op

# 1000个int数排序
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.0220 ns/op
BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.126 ns/op
```