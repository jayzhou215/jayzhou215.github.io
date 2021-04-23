package main

import (
	"fmt"
)

type A interface {
	Abs(a int) int
}

type T struct{}

func (t *T) Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var a A
	if a == nil {
		fmt.Printf("a is nil %v %T\n", a, a)
	}
	var b *T
	if b == nil {
		fmt.Printf("b is nil %v %T\n", b, b)
	}
	a = b
	if a == nil {
		fmt.Println("a hold b is nil")
	} else {
		fmt.Printf("a hold b is non nil %v %T\n", a, a)
	}
}
