package main

import "fmt"

func ListCase() {
	L := LinkList{}
	L.Init()
	for i := 1; i <= 5; i += 1 {
		L.PushBack(i)
	}
	L.Print()
	L.PushFont(6)
	L.Print()
	ok := L.Find(5)
	fmt.Println(ok)
	L.Insert(7, 8)
	L.Print()
	L.Insert(8, 2)
	L.Print()
	L.Update(5, 4)
	L.Print()
}
