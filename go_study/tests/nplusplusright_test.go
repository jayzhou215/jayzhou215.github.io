package tests

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Test_nPlusPlusRight(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "n++ right",
			want: 200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nPlusPlusRight(); got != tt.want {
				t.Errorf("nPlusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nPlusPlusRight2(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "n++ right",
			want: 200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nPlusPlusRight2(); got != tt.want {
				t.Errorf("nPlusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
