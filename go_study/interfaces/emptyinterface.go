package main

import "fmt"

func main() {
	var i interface{}
	describe_(i)

	i = 42
	describe_(i)

	i = "hello"
	describe_(i)
}

func describe_(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
