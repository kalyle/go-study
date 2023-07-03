package main

import "fmt"

// main 包默认不会加载其他文件， 而其他包都是默认加载的
// main包下多个文件，需go build/run . 进行编译运行
// func main() {

// }

func main() {
	fmt.Printf("地址为%p", &[3]int{1, 2, 3})
}
