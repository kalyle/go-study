package main

type Node struct {
	Val  int
	Next *Node
}

// type Node1 struct {
// 	val  int
// 	next *Node
// }
// 大写不报错，小写就会报错

func (list Node) PushBack(val int) bool {
	return true
}
func (list Node) PushFont(val int) bool {
	return true
}

func (list Node) Insert(val int, index int) bool {
	return true
}

func (list Node) Delete(index int) {

}
func (list Node) Update(val int) {

}
func (list Node) Find(val int) {

}
