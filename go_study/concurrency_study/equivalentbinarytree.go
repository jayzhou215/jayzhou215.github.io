package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	go Walk(t1, ch)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for {
		val, ok := <-ch
		val2, ok2 := <-ch2
		if (ok && !ok2) || (!ok && ok2) { // 两者数量不一致
			return false
		}
		if !ok && !ok2 { // 两者数量一致，且都等到了channel close()
			return true
		}
		fmt.Println(val, val2)
		if val == val2 {
			continue
		}
		return false
	}
}

func main() {
	//ch := make(chan int)
	//go Walk(tree.New(1), ch)
	//for val := range ch  {
	//	fmt.Println(val)
	//}
	t1 := tree.New(5)
	t2 := tree.New(5)
	ret := Same(t1, t2)
	if ret {
		fmt.Println("same", t1, t2)
	} else {
		fmt.Println("not same", t1, t2)
	}
}
