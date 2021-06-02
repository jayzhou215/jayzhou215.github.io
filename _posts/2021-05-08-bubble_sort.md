---
layout: post
title: Algorithm - bubble sort
tags: [go]
readtime: true
comments: true
---

### bubble sort
1. 冒泡排序的定义 它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢"浮"到数列的顶端。
2. 基本步骤
    1. 外循环是遍历每个元素，每次都放置好一个元素；　　　
    2. 内循环是比较相邻的两个元素，把大的元素交换到后面；
    3. 等到第一步中循环好了以后也就说明全部元素排序好了；
2. 示例代码分析
    1. 定义swapped变量，使用do while 跳出循环，对比循环n次，较少无效循环
    2. 定义indexToLastUnsortedElement，减少后续遍历中需要遍历的个数
    3. 匹配条件 left > right do swap
3. [bubble code](../algorithm/sort/sort_test.go)
4. 时间空间复杂度
    1. 时间复杂度: 外循环和内循环以及判断和交换元素的时间开销
        * 最优情况，内外循环+判断+无需交换，(n-1)，即O(n)。（助力，这里对应到具体实现代码是如此，默认的冒泡算法最优时间复杂度是O(n^2)）
        * 最差情况，内外循环+判断+全部交换，n*(n-1)/2，即O(n^2)
        * 平均情况，仍为O(n^2)
    2. 空间复杂度: 交换元素时那个临时变量所占的内存空间
        * golang中是隐式的，空间为2，即O(1)
    3. 稳定性 stability
        * 因为判断条件是left > right，等值时不会发生交换，总是稳定的    
### 引申知识点
1. bubble sort 在对降序排列的数组排序时明显优于对无序数组
    * 原因是branch predictior，简单来说就是cpu会先猜你会按上一次的if else/continue分支进行执行，并且把结果算出来，而如果猜错了预执行的结果会被丢弃，再重新执行，导致delay。
2. to test bubble sort - 性能测试和功能性测试
    1. 功能性测试，需要target array和destination array
    2. 性能测试，只需要target array
        1. 生成测试数组，随机生成n个数，并加入到数组中
        2. 执行排序算法，跑20s，看循环次数，算平均每次耗时，得到单次算法耗时。使用的待排序数组分无序、降序
    
### 参考资料
* [visual go bubble sort](https://visualgo.net/en/sorting)
* [时空复杂度分析](https://blog.csdn.net/YuZhiHui_No1/article/details/44339711)
* [happy coder bubble sort](https://www.happycoders.eu/algorithms/bubble-sort/)