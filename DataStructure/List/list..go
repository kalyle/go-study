package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}
type LinkList struct {
	Head *Node
	Len  int
}

func (L *LinkList) Init() {
	L.Head = &Node{}
	L.Len = 0
}
func (L *LinkList) PushBack(val int) {
	head := L.Head.Next
	if head == nil {
		L.Head.Next = &Node{Val: val}
	} else {
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &Node{Val: val}
	}
	L.Len++
}

func (L *LinkList) PushFont(val int) {
	head := L.Head
	next := head.Next
	head.Next = &Node{Val: val}
	head.Next.Next = next
	L.Len++
}

func (L *LinkList) Insert(val int, index int) {
	head := L.Head
	i := 0
	for ; head.Next != nil && i != index; i += 1 {
		head = head.Next
	}
	if head.Next != nil && i == index {
		next := head.Next
		head.Next = &Node{Val: val}
		head.Next.Next = next
	} else {
		head.Next = &Node{Val: val}
	}
	L.Len++

}

func (L *LinkList) Delete(index int) {
	if index > L.Len {
		return
	}
	for i, head := 0, L.Head; head.Next != nil; head, i = head.Next, i+1 {
		if i == index-1 {
			next := head.Next.Next
			head.Next = next
			break
		}
	}
	L.Len--
}
func (L LinkList) Update(index int, val int) {
	if index > L.Len {
		return
	}
	for i, head := 0, L.Head; head.Next != nil; head, i = head.Next, i+1 {
		if i == index {
			head.Val = val
			break
		}
	}
}
func (L LinkList) Find(val int) bool {
	if L.Len == 0 {
		return false
	}
	for head := L.Head.Next; head != nil; head = head.Next {
		if head.Val == val {
			return true
		}
	}
	return false
}
func (L LinkList) Print() {
	for head := L.Head.Next; head != nil; head = head.Next {
		fmt.Printf("%v\t", head.Val)
	}
	fmt.Println("长度为", L.Len)
}
