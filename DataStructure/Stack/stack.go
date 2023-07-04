package main

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Node
	Len int
}

func (s Stack) InitStack() *Stack {
	return &Stack{Node: Node{}, Len: 0}
}

func (s Stack) Push(e int) {
	if s.Len == 0 {
		s.Next = &Node{Val: e}
	} else {
		head := s.Next
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &Node{Val: e}
	}
}

func (s Stack) Pop() int {
	if s.Len == 0 {
		return INT_MIN
	}
	return INT_MAX
}
func (s Stack) Destory() {

}

func (s Stack) Clear() {

}

func (s Stack) Empty() bool {
	return true
}

func (s Stack) GetTop() int {
	return 1
}
