---
layout: post
title: Algorithm - radix sort
tags: [go]
readtime: true
comments: true
---

### radix sort
1. 基数排序的定义 是一种非比较型整数排序算法，其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较。。
2. 基本步骤
    1. 将所有待比较数值（正整数）统一为同样的数位长度，数位较短的数前面补零。
    2. 从最低位开始，依次进行一次排序。
3. [radix sort code](../algorithm/sort/sort_test.go)
4. 时间空间复杂度
    1. 时间复杂度:
        * 为O(n*k) 其中n是排序元素个数，k是数字位数。 
    2. 空间复杂度: 
        * O(1)

### 优点
避免了count sort中对空间依赖过大的问题，空间受所使用的bucket影响

### 缺点


### bench测试
[code](../algorithm/sort/sort_test.go)

```sh
// 10w, bucket size 19, k=32
BenchmarkQuick-12          	1000000000	         0.00864 ns/op
BenchmarkRandomQuick
BenchmarkRandomQuick-12    	1000000000	         0.0109 ns/op
BenchmarkMerge
BenchmarkMerge-12          	1000000000	         0.0139 ns/op
BenchmarkRadix
BenchmarkRadix-12          	1000000000	         0.0451 ns/op

// 100w, bucket size 19, k=32(int范围内)
BenchmarkQuick
BenchmarkQuick-12          	1000000000	         0.103 ns/op
BenchmarkRandomQuick
BenchmarkRandomQuick-12    	1000000000	         0.103 ns/op
BenchmarkMerge
BenchmarkMerge-12          	1000000000	         0.151 ns/op
BenchmarkRadix
BenchmarkRadix-12          	1000000000	         0.554 ns/op

// 100w, k=6
BenchmarkQuick
BenchmarkQuick-12          	1000000000	         0.103 ns/op
BenchmarkRandomQuick
BenchmarkRandomQuick-12    	1000000000	         0.105 ns/op
BenchmarkMerge
BenchmarkMerge-12          	1000000000	         0.152 ns/op
BenchmarkRadix
BenchmarkRadix-12          	1000000000	         0.212 ns/op

// 100w, k=4(最大10000)
BenchmarkQuick
BenchmarkQuick-12          	1000000000	         0.125 ns/op
BenchmarkRandomQuick
BenchmarkRandomQuick-12    	1000000000	         0.132 ns/op
BenchmarkMerge
BenchmarkMerge-12          	1000000000	         0.159 ns/op
BenchmarkRadix
BenchmarkRadix-12          	1000000000	         0.139 ns/op
```