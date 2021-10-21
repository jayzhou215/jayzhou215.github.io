package tests

import (
	"sync"
)

func nPlusPlus() int64 {
	var wg sync.WaitGroup
	var n int64
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			n++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			n++
		}
	}()
	wg.Wait()
	return n
}
