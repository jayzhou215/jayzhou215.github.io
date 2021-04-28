## 这个Video列举出的问题
[哔哩哔哩视频源](https://www.bilibili.com/video/BV1UJ411m7U1?from=search&seid=17329437087578237649)
> 前言: 2009年Go就面世了~
>
> this video is on 2012.
>
> 作者是 Rob Pike (can't sure)
>
> 完整的代码在[这里](https://github.com/adityamenon/Google-IO_2012_Go-Concurrency-Patterns)

1. Why concurrency
    1. use concurrency to describe the real world
2. What is concurrency
    1. concurrency is not parallelism, what's difference, 演讲者在heroku上的演讲
    2. talks on https://tinyurl.com/goconcnotpar
    3. 核心强调并发
3. a model for software construction
    1. easy to understand
    2. easy to use
    3. easy to reason about
    4. you don't need to be an expert!
4. History
    1. Hoare's CSP in 1978

5. Code Part, boring functions...
    1. the main thread won't wait for goroutine
    2. goroutine的初始堆栈很小，可以开启成百上千的routine
    3. goroutine可以根据需要扩展他的goroutine
    4. goroutine底层存在复用，但不需要开发者关注，只需要创建使用即可
    5. channel混合了通信和同步在一次操作中
    6. 带缓存的channel，buffered channel
    7. channel
        1. boring() do something internal and return the channel
        2. two boring wait each other
        3. in fan-in func, a new channel merge the two boring function, and so that they won't wait for each other
        4. 可以将一个channel的接收作为另一个channel的发送
    8. select
        1. balabala...     
6. system software
    1. 三个并发
    2. 增加超时机制
    3. 查询多个任务源取，最快的
7. More party tricks
    1. [Chatroulette toy](https://tinyurl.com/gochatroulette)
    2. [Load balancer](https://tinyurl.com/goloadbalancer)
    3. [Concurrent prime sieve](https://tinyurl.com/gosieve)
    4. [Concurrent power series(by Mcllroy)](https://tinyurl.com/gopowerseries)
8. don't overdo it
    1. `Sync` & `sync/atomic` supplies mutexes, condition variables, etc.
    2. 用`atomic`实现计数器
9. Q&A
    1. bala bala..
    2. a very cheap memory allocate and collect

## 个人感想
1. Go lang的目的是将复杂的事情收敛，让开发者更简单。   
