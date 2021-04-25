package main

import "fmt"

func fibonacci5(c, quit chan int) {
	x, y := 0, 1
	n := 0
	select {
	case c <- x:
		x, y = y, x+y
	case <-quit:
		fmt.Println("quit")
		return
	default:
		fmt.Println("default", n)
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 1; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci5(c, quit)
}
