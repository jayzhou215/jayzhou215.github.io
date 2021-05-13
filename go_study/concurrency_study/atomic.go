package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Val int
}

func main() {
	var data *Data
	go func() {
		for {
			if data == nil {
				data = &Data{5}
			}
		}
	}()
	go func() {
		for {
			if data != nil {
				data = nil
			}
		}
	}()

	go func() {
		for {
			if data != nil {
				if data.Val != 5 {
					// 这里不会命中
					fmt.Println("data.val not 5, current is", data.Val)
				}
			}
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
