package main

import (
	"fmt"
	"unsafe"
)

// 内存空间为0 数据结构 : struct{}, [0]T{}
type FS struct {
	E struct{}

	A uint32
	B uint64
	C uint64
	D uint64
}
type MS struct {
	A uint32
	E struct{}
	B uint64
	C uint64
	D uint64
}
type BS struct {
	A uint32
	B uint64
	C uint64
	D uint64
	E struct{}
}

// 结构体尾部size为0的变量(字段)会被分配内存空间进行填充，原因是如果不给它分配内存，该变量(字段)指针将指向一个非法的内存空间(类似C/C++的野指针)。

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof(FS{})) //32
	fmt.Println(unsafe.Sizeof(MS{})) //32
	fmt.Println(unsafe.Sizeof(BS{})) // 40

}
