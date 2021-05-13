package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Val int
}

func main() {
	var data interface{}
	go func() {
		for {
			data = &Data{5}
		}
	}()
	go func() {
		for {
			a := 10
			data = &a
		}
	}()

	go func() {
		for {
			if data != nil {
				d, ok := data.(*Data)
				if ok && d.Val != 5 {
					// 这里会频繁命中，打印 data.val not 5, current is 10
					fmt.Println("data.val not 5, current is", d.Val)
				}
				d2, ok2 := data.(*int)
				if ok2 && *d2 != 10 {
					// 这里会频繁命中，打印 *d2 not 10, current is 5
					fmt.Println("*d2 not 10, current is", *d2)
				}
			}
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
