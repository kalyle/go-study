package main

import (
	"fmt"
	_case "go-study/Generic/case"
)

// 泛型
func main() {
	a1, b1 := 1, 2
	a2, b2 := "abc", "def"
	c1 := _case.GetSumOfNum(a1, b1)
	fmt.Println("getSumOfNum c1 =", c1)
	c2 := _case.GetSumOfStr(a2, b2)
	fmt.Println("getSumOfStr c2 =", c2)

}

// gopath ,goroot含义
// mod,work
