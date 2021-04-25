package main

import "fmt"

func fibonacci4(c, quit chan int) {
	x, y := 0, 1
	n := 0
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default", n)
		}
		n++
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci4(c, quit)
}
