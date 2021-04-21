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
		fmt.Println("a is nil")
	}
	var b *T
	if b == nil {
		fmt.Println("b is nil")
	}
	a = b
	if a == nil {
		fmt.Println("a hold b is nil")
	} else {
		fmt.Println("a hold b is non nil")
	}
}
