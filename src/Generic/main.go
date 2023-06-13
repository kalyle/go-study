package main

import "fmt"
import "_case"

type I1 interface {
	int8
	int
}

func main() {
	a1, b1 := 1, 2
	a2, b2 := "abc", "def"
	c1 := _case.getSumOfNum(a1, b1)
	fmt.Println("getSumOfNum c1 =", c1)
	c2 := _case.getSumOfStr(a2, b2)
	fmt.Println("getSumOfStr c2 =", c2)

}
