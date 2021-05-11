> 关于atomic的使用，下列资料已经足够详细，不再赘述 

## 资料 -- actually, they are just google top3...
[Go 语言标准库中 atomic.Value 的前世今生](https://blog.betacat.io/post/golang-atomic-value-exploration/)
[golang atomic包的使用](https://www.jianshu.com/p/228c119a7d0e)
[Go语言——原子操作](https://www.jianshu.com/p/ccfbe7bf82bb)

## 要点
* atomic.Value{} 可以Load()或 Store()任意类型
* 对于任意类型 interface{} 在Store()方法中如何保证原子性的
    1. `StorePointer()`方法提供了对addr的原子性操作
    2. 任意类型的存储都可以拆分为对type, data的地址存储
    3. 故只要保证这两步是原子存储即可
        1. 当v.typ = nil时，使用`runtime_procPin` `runtime_procUnPin`来防止gc等
        2. 使用`^uintptr(0)`作为v.typ的中间态
        3. 使用`CompareAndSwapPointer()`来比较和交换addr，实现抢锁
        4. 当v.typ != nil时，直接使用`StorePointer()`来更新v.dat即可。
* unsafe.Pointer，直接操作内存，but，不保证向后兼容
* i++是非原子的，包括取值，加法，赋值
* `CompareAndSwapInt32()`的使用考虑配合循环

* [here](./concurrency_study/atomic.go) is just a very simple example that help to understand atomic.