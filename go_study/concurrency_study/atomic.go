package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	config := atomic.Value{}
	go func() {
		for {
			data := config.Load()
			if data != nil {
				themap, ok := data.(map[string]string)
				if ok {
					themap["hello"] = "world"
					return
				}
			}
		}
	}()
	config.Store(make(map[string]string))
	for {
		data := config.Load()
		if data != nil && len(data.(map[string]string)) > 0 {
			for key, value := range data.(map[string]string) {
				fmt.Println(key, value)
				return
			}

		}
	}

}
