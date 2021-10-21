package main

import (
	"fmt"
	"sync"
)

func main() {
	nPlusPlus()
}

func nPlusPlus() {
	var wg sync.WaitGroup
	var n int
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10*10000; i++ {
			n++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10*10000; i++ {
			n++
		}
	}()
	wg.Wait()
	fmt.Println(n)
}
