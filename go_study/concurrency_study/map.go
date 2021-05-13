package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//syncMap := sync.Map{}
	//syncMap.Store("hello", "world")
	//val, ok := syncMap.Load("hello")
	//if ok {
	//	fmt.Println(val)
	//}
	cnt := atomic.Value{}
	cnt.Store(0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt.Store(cnt.Load().(int) + 1)
		}()
	}
	wg.Wait()
	fmt.Println("cnt:", cnt)

}
