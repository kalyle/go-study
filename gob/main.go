package main

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int8
	Name string
}

// 以字节缓冲模拟网络传输的简单例子
func main() {
	// var network bytes.Buffer
	// enc := gob.NewEncoder(&network)

	// dec := gob.NewDecoder(&network)

	// err := enc.Encode(P{127, 127, 3, "zhangsan"})
	// if err != nil {
	// 	log.Fatal("编码异常")
	// }
	// err = enc.Encode(P{4, 5, 6, "lisi"})
	// if err != nil {
	// 	log.Fatal("编码异常")
	// }

	// var q Q
	// err = dec.Decode(&q)
	// if err != nil {
	// 	log.Fatal("解码异常")
	// }
	// fmt.Println("解码数据：", q.Name, *q.X, *q.Y)
	// err = dec.Decode(&q)
	// if err != nil {
	// 	log.Fatal("解码异常")
	// }
	// fmt.Println("解码数据：", q.Name, *q.X, *q.Y)
	// en2file()
	de2file()

}

// gob一个典型的用途是RPC的参数和结果
// 该实现为流中的每种数据类型编译自定义编解码器，并且在使用单个编码器传输值流时缓解编译成本，效率最高。
