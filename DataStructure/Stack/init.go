package main

import "fmt"

func StackWithArray() {
	stack := make([]int, 0)
	for i := 1; i <= 5; i += 1 {
		stack = append(stack, i)
	}
	top := stack[len(stack)-1]
	fmt.Println(stack, top)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}
