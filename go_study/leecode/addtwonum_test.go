package leecode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var testData = []struct {
	L1   int
	L2   int
	Dest int
}{
	{
		L1:   342,
		L2:   465,
		Dest: 807,
	},
	{
		L1:   0,
		L2:   0,
		Dest: 0,
	},
	{
		L1:   999999,
		L2:   9999,
		Dest: 1009998,
	},
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	cur := head
	inc := 0 // 单独标记进位
	for {
		add := inc
		if l1 != nil {
			add += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			add += l2.Val
			l2 = l2.Next
		}
		if add >= 10 {
			inc = 1
			add -= 10 // 修正 add结果
		} else {
			inc = 0 // 重置进位
		}
		// 处理头节点
		if cur == nil {
			cur = &ListNode{Val: add}
			head = cur
		} else {
			// 处理后续节点
			cur.Next = &ListNode{Val: add}
			cur = cur.Next
		}
		// 退出条件，考虑进位
		if l1 == nil && l2 == nil && inc == 0 {
			break
		}
	}
	return
}

func TestGetListNode(t *testing.T) {
	node := getListNode(345)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestGetNumFromListNode(t *testing.T) {
	node := getListNode(345)
	ret := getNumFromListNode(node)
	if ret == 345 {
		t.Log("ok")
	} else {
		t.Error(ret)
	}
}

//Runtime: 4 ms, faster than 98.55% of Go online submissions for Add Two Numbers.
//Memory Usage: 4.8 MB, less than 94.15% of Go online submissions for Add Two Numbers.
func TestAddTwoNumbers(t *testing.T) {
	for _, datum := range testData {
		l1 := getListNode(datum.L1)
		l2 := getListNode(datum.L2)
		ret := addTwoNumbers(l1, l2)
		numRet := getNumFromListNode(ret)
		if numRet == datum.Dest {
			t.Log(datum.Dest, "ok")
		} else {
			t.Error(datum.Dest, numRet)
		}
	}
}

func getNumFromListNode(node *ListNode) (ret int) {
	multi := 1
	for node != nil {
		ret += node.Val * multi
		node = node.Next
		multi *= 10
	}
	return
}

func getListNode(val int) (head *ListNode) {
	// 将数字342从低位到高位放到list中，最终head存2，依次存 4, 3
	cur := head
	for {
		mod := val % 10
		val = val / 10
		if cur == nil {
			head = &ListNode{Val: mod}
			cur = head
		} else {
			tmp := &ListNode{Val: mod}
			cur.Next = tmp
			cur = tmp
		}
		if val == 0 {
			break
		}
	}
	return
}
