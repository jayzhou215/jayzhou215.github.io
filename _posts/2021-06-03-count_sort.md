---
layout: post
title: Algorithm - count sort
tags: [go]
readtime: true
comments: true
---

### count sort
1. 计数排序的定义 是一种稳定的线性时间排序算法。
2. 基本步骤
    1. 找出待排序的数组中最大和最小的元素
    2. 统计数组中每个值为i的元素出现的次数，存入数组 C 的第i项
    3. 对所有的计数累加（从  C 中的第一个元素开始，每一项和前一项相加）
    4. 反向填充目标数组：将每个元素i放在新数组的第C[i]项，每放一个元素就将C[i]减去1
2. 示例代码分析
3. [count sort code](../algorithm/sort/sort_test.go)
4. 时间空间复杂度
    1. 时间复杂度:
        * 均为O(n+k)
    2. 空间复杂度: 
        * O(k)，k为数组中的最大值，可以用基数排序进行优化
    3. 稳定性 stability
        * unstable

### 优点
在数组中最大值较小时，非常快

### 缺点
在数组中最大值较大时，需要过多的空间和时间

### bench测试
[code](../algorithm/sort/sort_test.go)

```sh
// 10000 max 10000
BenchmarkBubble
BenchmarkBubble-12         	1000000000	         0.106 ns/op
BenchmarkSelect
BenchmarkSelect-12         	1000000000	         0.0510 ns/op
BenchmarkInsert
BenchmarkInsert-12         	1000000000	         0.0241 ns/op
BenchmarkQuick
BenchmarkQuick-12          	1000000000	         0.00125 ns/op
BenchmarkRandomQuick
BenchmarkRandomQuick-12    	1000000000	         0.000877 ns/op
BenchmarkMerge
BenchmarkMerge-12          	1000000000	         0.00233 ns/op
BenchmarkCount
BenchmarkCount-12          	1000000000	         0.000535 ns/op


// 10w max 10w
BenchmarkBubble
BenchmarkBubble-12         	       1	13543686470 ns/op
BenchmarkSelect
BenchmarkSelect-12         	       1	5221286440 ns/op
BenchmarkInsert
BenchmarkInsert-12         	       1	1971115760 ns/op
BenchmarkQuick
BenchmarkQuick-12          	1000000000	         0.0105 ns/op
BenchmarkRandomQuick
BenchmarkRandomQuick-12    	1000000000	         0.0118 ns/op
BenchmarkMerge
BenchmarkMerge-12          	1000000000	         0.0144 ns/op
BenchmarkCount
BenchmarkCount-12          	1000000000	         0.00487 ns/op
```