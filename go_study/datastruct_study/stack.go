package datastruct_study

import (
	"fmt"
	"math/rand"
)

const maxN = 1000

type Node struct {
	Val  int
	Next *Node
}

type Stack interface {
	Create(n int)
	Peek() *Node
	Push(*Node)
	Pop() *Node
}

type StackImplementation1 struct {
	NodeList []*Node
}

func (s *StackImplementation1) Create(n int) {
	s.NodeList = make([]*Node, n)
	for i := 0; i < n; i++ {
		s.NodeList[i] = &Node{Val: rand.Intn(maxN)}
		fmt.Println("create", i, s.NodeList[i].Val)
		if i > 0 {
			s.NodeList[i-1].Next = s.NodeList[i]
		}
	}
	return
}

func (s *StackImplementation1) Peek() *Node {
	return s.NodeList[len(s.NodeList)-1]
}

func (s *StackImplementation1) Push(node *Node) {
	s.NodeList = append(s.NodeList, node)
}

func (s *StackImplementation1) Pop() *Node {
	if len(s.NodeList) == 0 {
		return nil
	}
	tail := s.NodeList[len(s.NodeList)-1]
	s.NodeList = s.NodeList[:len(s.NodeList)-1]
	return tail
}

type StackImplementation2 struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (s *StackImplementation2) Create(n int) {
	var preNode *Node
	for i := 0; i < n; i++ {
		node := &Node{
			Val: rand.Intn(maxN),
		}
		fmt.Println("create", i, node.Val)
		if i == 0 {
			s.Head = node
			preNode = node
		}
		if i > 0 {
			preNode.Next = node
			preNode = node
		}
		if i == n-1 {
			s.Tail = node
		}
		s.Length++
	}
	return
}

func (s *StackImplementation2) Peek() *Node {
	return s.Tail
}

func (s *StackImplementation2) Push(node *Node) {
	if s.Tail != nil {
		s.Tail.Next = node
		s.Tail = node
	} else {
		s.Tail = node
		s.Head = node
	}
}

func (s *StackImplementation2) Pop() *Node {
	var bef *Node = s.Head
	for bef != nil && bef.Next != nil && bef.Next != s.Tail {
		bef = bef.Next
	}
	last := s.Tail
	if bef != nil {
		bef.Next = nil
	}
	s.Tail = bef
	return last
}
