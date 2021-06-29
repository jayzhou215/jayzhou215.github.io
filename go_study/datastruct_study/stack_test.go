package datastruct_study

import (
	"fmt"
	"testing"
)

func TestStackIm1(t *testing.T) {
	si := StackImplementation1{}
	si.Create(10)
	for i := 0; i < 10; i++ {
		pop := si.Pop()
		if pop != nil {
			fmt.Println(i, pop.Val)
		}
	}
}

func TestStackIm2(t *testing.T) {
	si := StackImplementation2{}
	si.Create(10)
	for i := 0; i < 10; i++ {
		pop := si.Pop()
		if pop != nil {
			fmt.Println(i, pop.Val)
		}
	}
}
