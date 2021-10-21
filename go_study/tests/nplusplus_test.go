package tests

import (
	"sync"
	"testing"
)

func Test_nPlusPlus(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "n++",
			want: 200000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nPlusPlus(); got != tt.want {
				t.Errorf("nPlusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
