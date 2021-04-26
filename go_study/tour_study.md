[toc]
## [官方教程](https://tour.golang.org/welcome/1) 学习笔记
#### 基础
1. 第一个Hello, world ~~ 
  * 右侧的编辑器支持语法高亮
2. Go本地化，尝试挑战下英文
3. 离线执行该教程
  * 在命令行将 `go get golang.org/x/tour` 执行完之后， 直接敲 `tour`，会在浏览器打开http://127.0.0.1:3999/welcome/1
  * [go get命令](http://c.biancheng.net/view/123.html) 

#### 包、变量、函数
1. 包名是路径的最后一个元素
2. 包的导入可以分成多行，不过分组导入语句是好的形式
3. 只有变量或方法的首字母是大写时，才是已导出的，可以在包外使用，否则只能在包内使用
4. 函数 - 类型在变量名之后
5. 形参是相同类型时，可以省略中间的
6. 函数可以返回多个值
7. 可以命名返回值
8. 变量 - var 可以声明变量列表
9. 变量声明可以包含初始值，形如 `var x, y, z = true, false, "no"`
10. 短变量声明 `:=`，仅用于函数内
11. 变量类型，特别的`rune, uintptr`，拓展阅读1 [2](https://golangbyexample.com/understanding-uintptr-golang/)
12. 零值，没有明确声明的都会赋予零值 int 0, bool false, string ""
13. 类型转换 `T(v)`
14. 类型推导，没有明确声明的由右式推导得出
15. 常量，使用const声明，不能用 `:=`
16. 数值常量是高精度的值 `const Big = 1 << 100`，示例中int可能是32位，可能是64位，但都比100位小

#### 流控制语句 for, if else, switch and defer
1. for 没有小括号，有大括号
2. 初始、后置可选
3. 可以去掉`;`，替代`while`
4. 无限循环 `for{}`
5. if 没有小括号，有大括号
6. if 简短语句 `if v:= a+b; v < lim {}`
7. `if {} else {}`
8. 练习: 循环与函数，翻译的叫牛顿法，其实就是算一个误差，反复逼近。需要假定一个结果值，计算误差，用误差再去修正假定值。
```go
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=0;i<=7; i++{
		z -= (z*z - x) / (2*z)
	}
	return z
}


func main() {
	fmt.Println(Sqrt(10))
}
```
9. switch，go中的case自带break
10. 顺序执行，匹配成功则退出，case后跟的可以是func()
11. 没有条件的switch = switch true {}
12. defer
    1. 执行时机推迟到外层函数返回之后
    2. 推迟调用的函数的参数会立即求值
13. defer执行栈，后进先出

#### 更多类型: struct, slice and map
1. 指针
    1. 指针保存了值的内存地址
    2. `var p *int` 零值为 `nil`
    3. `i := 42` `p = &i`
    4. `fmt.Println(*p)` // 取值
    5. `*p = 21` // 设值
2. 结构体，就是一组字段
    ```go
    package main
    
    import "fmt"
    
    type Vertx struct {
        X int
        Y int
    }
    
    func main() {
        fmt.Println(Vertx{1, 2})
    }
    
    ```
3. 取值用点号 `vertx.X = 4`
4. 结构体指针，`(*p).X`  `p.X` 第二种进行了隐式引用
5. 结构体文法，直接列出值来新分配一个结构体 `v1 = Vertx{1, 2}`
6. 数组，`[n]T` 表示n个T类型的数组 `var a [10]int` 
    1. 数组大小是固定的
    2. 数组声明后可以直接使用 `a[0] = "Hello" a[1] = "world"`
7. 切片，`[]T`表示一个类型为T的切片
    1. a[low : high] 前包后闭， a[1:4]取的是位置1到3的元素
8. 切片并不存储任何数据，只是描述了底层数组中的一段。
    1. 更改切片会修改底层数组中对应的元素。
    2. 与他共享底层数组的切片都会观测到这些修改。
9. 数组语法vs切片语法 `[3]bool{true, true, false}` vs `[]bool{true, true, false}`
10. 切片默认行为 `var a [10]int` `a[:10]` `a[0:]` `a[0:10]` `a[:]`
11. 切片的长度和容量
    1. 长度就是它所包含的元素个数。
    2. 容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
    3. 切片的上界不能超过切片容量
12. 切片的零值
    1. 切片的零值是nil
    2. 零值切片`var a []int`的`len(a) = 0` `cap(a) = 0`
13. 用make创建切片
    1. `a := make([]int, 5)`  // len(a)=5
    2. `b := make([]int, 0, 5)` // len(b)=0, cap(b)=5
    3. `b = b[:cap(b)]` // len(b)=5, cap(b)=5
    4. `b = b[1:]`      // len(b)=4, cap(b)=4
14. 切片的切片
    1. 切片可包含任何类型，甚至包括其它的切片。
15. 向切片追加元素
    1. `func append(s []T, vs ...T) []T
`   2. append中的切片可以是nil
    3. 接收返回值作为变更后的切片
    4. append过程中，内部会自动扩充底层数组
    5. `slice = append(slice, elem1, elem2)`
    6. `slice = append(slice, anotherSlice...)`
    7. `slice = append([]byte("hello "), "world"...)`
    8. [引申阅读](https://blog.go-zh.org/go-slices-usage-and-internals)
16. Range
    1. for 循环的 range 形式可遍历切片或映射。
    2. 当使用 for 循环遍历切片时，每次迭代都会返回两个值。
    3. 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份**副本**。
    4. 指定切片类型为struct的指针类型，来避免struct本身的copy
17. Range(续)
    1. `for i, _ := range pow`
    2. `for _, value := range pow`
    3. `for i := range pow`
18. slice exercise
    1. slice的构造方式有两种，一种是make，一种是声明后append()
    2. 计算结果需要做类型转换
19. Map
    1. 使用make初始化 `make(map[string]Vertex)`
    2. key value存储
20. Map初始化
    ```go
    var m = map[string]Vertex{
        "Bell Labs": Vertex{
            40.68433, -74.39967,
        },
        "Google": Vertex{
            37.42202, -122.08408,
        },
    }
    ```
21. Map初始化续 子类型可以忽略
    ```go
    var m = map[string]Vertex{
        "Bell Labs": {40.68433, -74.39967},
        "Google":    {37.42202, -122.08408},
    }
    ``` 
22. 操作Map
    ```go
    m[key] = elem // 增更
    elem = m[key] // 取值
    delete(m, key) // 删除
    elem, ok := m[key] // 测试一个key是否已经存在
    ```
23. [Map练习](./wordcount/main.go)
    1. string直接遍历时取出的是rune
    2. 句中的空格分隔符可能出现在开始、中间、末尾
24. function类型的值
    1. function也是一种值类型，可以被传递
    2. 可以作为方法入参，也可以作为返回值
25. function闭包
    1. go的function也可以是一种闭包
    2. 闭包是一个function值，在他的方法体外引用变量
    3. 疑问: 闭包的使用场景是什么
        1. https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
            1. isolate data
                1. [fibgen](./functionclosures/fibgen.go)
                    1. 斐波那契数列生成，持有f1, f2, 又动态变更其值
                2. [maze](./functionclosures/maze.go)，想到一道面试题，用闭包来实现很简单，需求是在一个二维数组中按顺时针由外环向内环找到第n个数
                3. 这种使用场景和带有成员变量的类很像，当然在go里面是一个struct，自带function
            2. Wrapping functions and creating middleware
                1. [timed middle ware](./functionclosures/timedmiddleware.go)
            3. Accessing data that typically isn’t available
                1. [处于对数据保护的考虑，不想向外暴露特定字段时使用](./functionclosures/varprotection.go)
            4. [Binary searching with the sort package](./functionclosures/search.go)
                1. `sort.SearchInts()`使用闭包
                2. todo 对sort package中提供的各种排序、搜索算法进行梳理学习
            5. Deferring work 推迟执行
            ```go
            go func() {
             result := doWork1(a, b)
             result = doWork2(result)
             result = doWork3(result)
             // Use the final result
           }()
           fmt.Println("hi!")
            ```

#### Methods and Interfaces
1. Methods
    1. Go不存在class，但是可以在类型上定义func
    2. 一个 Method 是一个带有特定receiver的function
    3. receiver出现在func关键字和method name之间
2. 一个Method只是一个带有receiver参数的function，可以将receiver参数写到参数列表中
3. 也可以在非struct类型上声明method
    1. 只能在同一个包为一个类型声明带receiver的method
4. 指针receiver
    1. 可以声明指针 receiver的method
    2. receiver类型是星语法
    3. 指针 receiver比值 receiver更常用，因为经常需要在method里面修改 receiver
    4. 值receiver是receiver的一个copy，修改值receiver并不修改原始值
5. 指针和方法
    1. 对比值引用和指针引用的区别
6. Method和指针的间接引用
    1. 类型是方法入参时，如类型是指针类型，调用时所传入参必须是指针类型。
    2. 类型是receiver参数时，method的receiver类型是指针类型时，调用时receiver类型是值类型或者指针类型都可以，Go语言内部会做转换
7. 反之亦然
    1. 类型是方法入参时，如类型是值类型，调用时所传入参必须是值类型。                
    2. 类型是receiver参数时，method的receiver类型是值类型时，调用时receiver类型是值类型或者指针类型都可以，Go语言内部会做转换
8. 选择值还是指针receiver
    1. 有两个原因选择指针receiver
        1. method可以修改指针receiver所指向的值
        2. 避免在method调用时的值copy，这在receiver所指向的struct是个大对象时会很有用
    2. 一般不会混用值或者指针receiver
9. interface
    1. 一个interface类型定义是一组method的签名(方法名称，入参、出参类型，并不包含func关键字)
    2. 一个interface类型的值可以持有任何实现了那些method的值
    3. interface的值必须是指针类型，而不允许是值类型，这里没有隐式转换
10. interface是被隐式实现
    1. 实现interface无需显示声明，没有`impletement`关键字
    2. 隐式interface连接interface定义及其实现，可以无需其他预先的设置就可以在package的任何地方使用
11. interface values
    1. 实际上，interface values可以被认为是一个tuple，由值和具体类型`value type`
    2. interface value持有一个特定底层具体类型的value
    3. 调用一个基于一个interface的value的method是执行基于它的底层类型的同名method
12. 底层是nil value的interface values
    1. 如果在interface内部的具体值是nil，这个method会通过一个nil receiver被调用
    2. 在某些语言这会触发一个空指针异常，但是在Go，写method的时候很友善的处理被一个nil receiver调用的情况的是很常见的
    3. 注意，[一个持有nil的实际值的interface值本身是非nil的](./interfaces/main.go)
13. [nil interface values](./interfaces/nilinerface.go)
    1. 一个nil interface value既不持有值也不持有其具体类型
    2. 在一个nil interface上调用一个method会触发一个运行时异常，因为在interface tuple中没有类型，其来指明调用哪个具体的method
14. [empty interface](./interfaces/emptyinterface.go)
    1. 指定0个methods的 interface类型被称为 empty interface
    2. 一个空的empty interface可以持有任何类型的值
    3. 空interface 被那些处理未知类型的代码使用。比如`fmt.Print()`接收任何数量的`interface{}`类型    
15. 类型断言
    1. 一个类型断言提供一个interface 值的底层具体值的入口
    2. `t = i.(T)` 这种声明断言interface i持有具体的类型 T，并将底层的类型T赋给变量t
    3. 如果i并不持有类型T，这个语句会触发panic
    4. `t, ok = i.(T)`为了测试一个interface value是否持有特定类型T，一个类型断言可以返回两个值，具体类型值和一个报告断言是否成功的bool值
    5. 如果i持有T，t是T的具体类型值，ok是true，反之，t是T的零值，ok是false，并且没有panic
    6. 注意和从map中取值的语法相似性
16. type switch
    1. 一个type switch是一个允许多个有序类型断言的结构
    2. 一个type switch像一个普通的switch语句，但是在一个type switch的cases指定的是type而不是value，并且这些values和给定的interface所持有的类型进行比较
    3. 一个type switch的声明与一个类型断言`v = i.(T)`有相同语法，但是特定类型T被关键字`type`替换
    4. i持有类型T或S，在各个T/S的case语句中，变量v持有对应的T或S类型，且持有变量i所持有的value，在default case中v持有和i相同的value和type
17. Stringers
    1. 最普遍的interface就是在fmt包中定义的Stringers interface
    2. Stringers是一个将自己描述为一个字符串的类型。fmt包寻找这个interface来打印值
18. Stringers exercise
    1. 实现String() string方法
    2. 注意fmt.Printf()这里没有隐式转换，示例中的打印的ip类型是IPAddr，则会去找receiver类型为值的String()方法。原因是底层使用的switch type方式去判断类型
    3. 换句话说，实现了String()方法之后，还需要注意在fmt中传值的类型和String()方法定义的receiver类型一致
    ```
    switch v := p.arg.(type) {
    case error:
        handled = true
        defer p.catchPanic(p.arg, verb, "Error")
        p.fmtString(v.Error(), verb)
        return

    case Stringer:
        handled = true
        defer p.catchPanic(p.arg, verb, "String")
        p.fmtString(v.String(), verb)
        return
    }
    ```
19. Go程序通过`error`值表达error状态
    1. error type是一个内置interface，和fmt.Stringer类型
    2. 和fmt.Stinger一样，fmt包在打印信息时会寻找error interface
    3. Functions经常返回error value，调用代码需要通过测试err是否为nil来处理errors
    4. A nil error 表示成功，非nil error表示失败
20. [error exercise](./interfaces/errors.go)
    1. [算法原理](https://zhuanlan.zhihu.com/p/58754724)
    2. 算法解释: 
        1. 假定一个值z
        2. 假定值z*z - x得到误差
        3. 按照 `2*z` 的速率计算得到一个误差修正值 -- 如果不用`2*z`而是用`3*z`或者`1.5*z`只是会慢些逼近。`2*z`在数学上也是z的平方的导数，用导数去逼近会更快。
        4. 用这个修正值去修正z
    3. 值为-2时，不做错误检查，结果是0.5353752385394379
21. Readers
    1. io包指定了io.Reader interface，代表了数据流读的结尾
    2. Go中提供的基础包 files, network connections, compressor, cipher...
    3. Read弹出给定byte切片数据并且返回弹出的字节数和一个error值
    4. 当stream end的时候返回一个io.EOF error
    5. 示例中`b[:n]` 如果把 `[:n]`去掉，可以发现上次读到的数据还在其中
22. [Read exercise](./interfaces/reader.go)
23. [rot13 exercise](./interfaces/rot.go)
    1. 整个完整的执行链路
        1. 查看`io.Copy()`源码，确定在这里会执行这段`nr, er := src.Read(buf)`，其中`buf`是重新创建的
        2. Read()方法的链式调用
            1. `"Lbh penpxrq gur pbqr!"`这个是需要加密的字符串
            2. `string.NewReader()`中的`Read()`会去`r.s`中去读取byte
            3. 在`rot13.Reader()`中，需要先去执行`io.Reader`的`Read()`方法，拿到原始输入
            4. 针对原始输入进行rot13加密
    2. [rot13的原理](https://en.wikipedia.org/wiki/ROT13)
24. Images
    1. 三个methods的interface
25. [Image exercise](./interfaces/images.go)
    1. can try `(x+y)/2, x*y, and x^y` here, like in [slice exercise](http://127.0.0.1:3999/moretypes/18)

#### Concurrency
1. Goroutines
    1. 一个go routine是一个被Go运行时管理的轻量级线程
    2. `go f(x, y, z)` 开始一个运行中的goroutine
    3. x,y,z的赋值在当前线程，f的执行在新的线程中
    4. Go routine运行在相同的地址空间中，所以对共享内存的处理必须是同步的
    5. sync包提供了有用的原始工具，尽管你不会需要这些，因为有其他的工具。
2. Channels
    1. channel是一个特定类型的管道，通过它，你可以收发value，通过channel operator `<-`
    2. 像map, slice，channel需要先创建再使用 `ch := make(chan int)`
    3. `ch <- value` 将value传递给channel， `v := <- ch` 从channel中取值并传递给变量v，按照箭头方向的数据流
    4. 默认情况，senders和receivers会block，直到另一侧ready。这种特性，允许协程同步而无需直观锁或者条件变量
    5. example中
        1. goroutine中赋值，在主线程中取值
        2. ch被两个线程赋值，在主线程中取值了两次，这意味着在不同goroutine可以考虑使用同一个channel // surprise me
3. Buffered Channels
    1. Channels 可以被 buffered，在make方法的第二个参数中提供 buffered length用`make`来初始化一个buffered channel
    2. 发送一个buffered channel仅当channel是满的时候会block。当buffer是空的时候，接收者会block。
    3. 修改示例装满buffer来看会发生什么
        1. [`ch := make(chan int)`的行为](./concurrency_study/channel1.go)和[`ch := make(chan int, 1)`的行为](./concurrency_study/channel2.go)是不一致的
        2. 前者，在同一线程中，一旦执行收/发即会lock，需要在其他线程执行相反动作来unlock
        2. 后者，需要超过声明中的数量，才会lock
4. Range and Close
    1. 一个sender可以`close`一个channel来指示没有更多的值会被发送。接收者在接收者表达式中通过分配第二个参数来测试一个channel是否被关闭，`v, ok := <- ch`
    2. 当没有更多值接收并且channel被关闭时，`ok`值是`false`
    3. loop `for i := range c` 会从channel里面重复的接收value直到它关闭
        1. `range c` 循环的次数并不限定于 `channel`的定义数，而是发送数据到c的次数
        2. 需要主动在sender里面close channel，否则会deadlock
    4. 只能在sender里面发close，在receiver中发close会panic
    5. channel不像文件，一般不需要关闭它，close()仅当在一些必须要告诉接收者没有更多数据会来了的情况，比如为了关闭一个range loop
5. Select
    1. Select语法让一个routine在多个通信操作上等待
    2. 一个select阻塞直到其中一个case可以执行，然后它会执行那个case，如果多个同时ready，它会随机选个
        1. 这里自测发现矛盾的点，加上default，发现中间会穿插很多default行为
        2. [如果去掉for循环，只执行一次default，并不会阻塞](./concurrency_study/select2.go)
        3. 结论
            1. 仅select并不阻塞
            2. for + select组合时，由于没有default行为，看上去直到其中某一个case满足条件时才执行那个case
            3. 相当于一直在死循环，直到某个channel接收到值
6. Default Select
    1. 如5中所想，实际会走default，Good example!
    2. time.Tick()和time.After()返回的都是channel，发送方在time内部
7. 用channel来实现等价二叉树练习
    1. 可以有很多种所存储的是同样的值的二叉树
    2. 在大部分语言中匹配两个二叉树中的值是否一致是件很困难的事情
    3. 我们将用Concurrency和channels来实现一种简单写法
8. [练习](./concurrency_study/equivalentbinarytree.go) 
    1. tree.New()返回的是一个有序二叉树，值相同，结构不同
    2. 二叉树的遍历，应当从最下、最左子树开始，依次根节点，右节点
    3. Same中比较两个tree的Node值，如果得到需要遍历的N，或者如何终止遍历
        1. 使用闭包，实现channel的close() 参考[这里](https://stackoverflow.com/questions/12224042/go-tour-exercise-equivalent-binary-trees)   
    4. 虽然tree.New()总是返回10个值，如果需要考虑两者数量一致
9. Sync.Mutex 互斥
    1. 我们已经看到channel在goroutine间通讯是如何棒的
    2. 但是，如果我们不需要通讯呢，如果我们仅仅是想保证同时只有一个goroutine可以处理某个变量来避免冲突
    3. 这个概念叫做相互排斥(mutual exclusion)，提供这个功能的数据结构简称mutex，通过`sync.Mutex`及其方法`lock`和`unlock`        
    4. 我们可以定义一个代码块，被执行在互相排斥中，通过用一个lock及unlock包裹如Inc方法中所展示的
    5. 用defer unlock()来确保释放锁
10. web-crawler
    1. 并发使用`go`开启goroutine
    2. 避免重复请求
        1. 使用map
        2. 由于使用多线程，使用Lock, UnLock来避免冲突；仅对map操作部分进行lock,unlock，避免过多占用锁
        3. 这个功能点的hook入口，考虑新增一个crawler struct，收敛wg & mutex
    3. 开启goroutine后主线程需要等待routine执行完毕，`waitGroup` 总是 `Add(), Done(), Wait()`一起出现
    4. for循环中声明的临时变量u，如果在go func(){}()中如果使用了，需要注意做额外的传值处理，如示例
11. Where to Go from here...
    1. [Go Documentation](https://golang.org/doc/)
    2. Write go code
        1. [writing go code video(哔哩哔哩源)](https://www.bilibili.com/video/BV1pt41157WA?from=search&seid=15425332049712931967)
        2. [How to write Go code](https://golang.org/doc/code)
    3. Language
        1. [package reference](https://golang.org/pkg/)
        2. [Language specification](https://golang.org/ref/spec)
    4. Concurrent
        1. [Go Concurrency Patterns(哔哩哔哩源)](https://www.bilibili.com/video/BV1UJ411m7U1?from=search&seid=17329437087578237649)
        2. [Advanced Go currency Patterns](https://www.bilibili.com/video/BV177411A7F2?from=search&seid=2366347770921386959)
        2. [concurrent code](https://golang.org/doc/codewalk/sharemem/)
    5. Web application
        1. [writing web application](https://golang.org/doc/articles/wiki/)
    6. [Go function types](https://golang.org/doc/codewalk/functions/)
    7. [Go blog](https://blog.golang.org/)    