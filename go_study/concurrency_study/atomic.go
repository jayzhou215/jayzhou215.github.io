package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var cnt int64
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&cnt, 1)
		}()
	}
	wg.Wait()
	fmt.Println("cnt:", cnt)
}
