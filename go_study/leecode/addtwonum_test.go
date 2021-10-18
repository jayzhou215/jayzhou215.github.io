package leecode

import (
	"fmt"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 难点在于需要记录头节点，当前运行时节点，也可以叫尾节点
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	return handleNext(l1, l2, 0)
}
func handleNext(l1 *ListNode, l2 *ListNode, plus int) *ListNode {
	if l1 == nil && l2 == nil {
		if plus == 0 {
			return nil
		}
		return &ListNode{Val: 1}
	}
	l1Val := 0
	l2Val := 0
	var l1Next *ListNode
	var l2Next *ListNode
	if l1 != nil {
		l1Val = l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		l2Val = l2.Val
		l2Next = l2.Next
	}
	var node ListNode
	sum := l1Val + l2Val + plus
	if sum >= 10 {
		node.Val = sum - 10
		plus = 1
	} else {
		node.Val = sum
		plus = 0
	}
	node.Next = handleNext(l1Next, l2Next, plus)
	return &node
}

func TestGetListNode(t *testing.T) {
	node := getListNode([]int{3, 4, 5})
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestGetNumFromListNode(t *testing.T) {
	origin := []int{7, 0, 8, 9}
	node := getListNode(origin)
	ret := getNumFromListNode(node)
	if reflect.DeepEqual(ret, origin) {
		t.Log("ok")
	} else {
		t.Error(ret)
	}
}

//Runtime: 4 ms, faster than 98.55% of Go online submissions for Add Two Numbers.
//Memory Usage: 4.8 MB, less than 94.15% of Go online submissions for Add Two Numbers.
func TestAddTwoNumbers(t *testing.T) {

	var testData = []struct {
		L1   []int
		L2   []int
		Dest []int
	}{
		{
			L1:   []int{3, 4, 2},
			L2:   []int{4, 6, 5},
			Dest: []int{7, 0, 8},
		},
		{
			L1:   []int{3, 4, 2},
			L2:   []int{4, 6, 5, 9},
			Dest: []int{7, 0, 8, 9},
		},
		{
			L1:   []int{0},
			L2:   []int{0},
			Dest: []int{0},
		},
		{
			L1:   []int{9, 9, 9, 9, 9, 9},
			L2:   []int{9, 9, 9, 9},
			Dest: []int{8, 9, 9, 9, 0, 0, 1},
		},
	}

	for _, datum := range testData {
		l1 := getListNode(datum.L1)
		l2 := getListNode(datum.L2)
		ret := addTwoNumbers(l1, l2)
		numRet := getNumFromListNode(ret)
		if reflect.DeepEqual(numRet, datum.Dest) {
			t.Log(datum.Dest, "ok")
		} else {
			t.Error(datum.Dest, numRet)
		}
	}
}

func getNumFromListNode(node *ListNode) (ret []int) {
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	return
}

func getListNode(val []int) (head *ListNode) {
	// val [3,4,5], 数字543的逆序存放，最终 3->4->5
	head = &ListNode{}
	head.Val = val[0]
	if len(val) > 1 {
		head.Next = getListNode(val[1:])
	}
	return
}
