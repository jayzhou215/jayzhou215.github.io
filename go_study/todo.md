1. [slice 引申阅读](https://blog.go-zh.org/go-slices-usage-and-internals)
2. 对sort package中提供的各种排序、搜索算法进行梳理学习
3. [gophers exercise](https://gophercises.com/)
4. 梳理go中会出现隐式转换的地方
    1. 通过receiver调用方法时，receiver类型定义中无论是值还是指针，实际调用时无论是值还是指针都会隐式转换
    2. function中入参是interface{}时，入参的是值类型或指针类型均可(why?)
    3. 变量类型是interface时，给该变量赋值时也必须传指针类型，此处不会有隐式转换
     