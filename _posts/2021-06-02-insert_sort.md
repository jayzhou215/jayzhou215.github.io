---
layout: post
title: Algorithm - insert sort
tags: [go]
readtime: true
comments: true
---

### insert sort
1. 插入排序的定义 通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
2. 基本步骤
    1. 从第一个元素开始，该元素可以认为已经被排序
    2. 取出下一个元素，在已经排序的元素序列中从后向前扫描
    3. 如果该元素（已排序）大于新元素，将该元素移到下一位置
    4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
    5. 将新元素插入到该位置后
    6. 重复步骤2~5
2. 示例代码分析
    1. 循环n-1次
    2. 单次循环可以理解为 对有序序列进行一次反向的冒泡排序
3. [insert sort code](../algorithm/sort/sort_test.go)
4. 时间空间复杂度
    1. 时间复杂度:
        * 比较操作，最优n-1次，最坏n*(n-1)/2，即O(n^2)
        * 交换操作，最优0次，最坏n*(n-1)/2次，即O(n^2)
    2. 空间复杂度:
        * 交互在golang中是隐式的，空间为2，即O(1)
    3. 稳定性 stability
        * 可控，因为判断条件是left > right，等值时不会发生交换，总是稳定的    

### 优点
简单 & in-place(原地操作) & 对有序序列进行操作，减少了无效比较，且有优化空间

### 缺点
平均时间复杂度依然是O(n^2)

### bench测试
```sh
# 1000个Int数
BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.000879 ns/op
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.000732 ns/op
BenchmarkInsert
BenchmarkInsert-12    	1000000000	         0.000288 ns/op

# 10000个int数
BenchmarkBubble
BenchmarkBubble-12    	1000000000	         0.101 ns/op
BenchmarkSelect
BenchmarkSelect-12    	1000000000	         0.0513 ns/op
BenchmarkInsert
BenchmarkInsert-12    	1000000000	         0.0235 ns/op
```