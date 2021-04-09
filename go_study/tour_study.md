## [官方教程](https://tour.golang.org/welcome/1) 学习笔记
#### 基础
1. 第一个Hello, world ~~ 
  * 右侧的编辑器支持语法高亮
2. Go本地化，尝试挑战下英文
3. Go offline
  * `go get golang.org/x/tour` 执行完之后， 直接敲 `tour`，会在浏览器打开http://127.0.0.1:3999/welcome/1
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
9. 