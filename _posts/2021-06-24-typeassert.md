---
layout: post
title: go - why we can't invoke a method with pointer receiver from a type assertion struct
tags: [go]
readtime: true
comments: true
---

### here are examples that work and not work 
ok
```go

type Integer int

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func (a *Integer) Sub(b Integer) Integer {
	return *a - b
}

func TypeAssert1() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	sub := i.(*Integer).Sub(b)
	fmt.Println(sum, sub)
}

``` 

ok
```go

type Integer int

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func (a *Integer) Sub(b Integer) Integer {
	return *a - b
}


func TypeAssert1() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = a
	sum := i.(Integer).Add(b)
    sub := i.(Integer).Sub(b) // compile error  cannot call pointer method on i.(Integer) cannot take the address of i.(Integer)
	fmt.Println(sum, sub)
}

``` 

## the reason is 
We can't (in this case implicitly for a pointer receiver) take the address of the result of an expression (b.(Integer)). We can take the address of a variable. 

## references
https://stackoverflow.com/questions/43883502/how-to-invoke-a-method-with-pointer-receiver-after-type-assertion