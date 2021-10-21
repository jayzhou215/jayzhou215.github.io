package tests

import (
	"sync"
	"sync/atomic"
)

func nPlusPlusRight() int64 {
	var wg sync.WaitGroup
	var n int64
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			atomic.AddInt64(&n, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			atomic.AddInt64(&n, 1)
		}
	}()
	wg.Wait()
	return n
}

func nPlusPlusRight2() int64 {
	var wg sync.WaitGroup
	var n int64
	wg.Add(2)
	mu := sync.Mutex{}
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()
	wg.Wait()
	return n
}
