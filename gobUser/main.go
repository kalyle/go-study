package main

import (
	"bytes"
	"encoding/gob" //仅用于go语言之间编码
	"fmt"
	"log"
)

type Stu struct {
	Id   int
	Name string
}
type Pel struct {
	Id   int
	Name string
}

func main() {
	// defer func(err string){
	// 		fmt.Println("捕获异常:",err)
	// 	}(err)
	var network bytes.Buffer            //标准输入
	encoder := gob.NewEncoder(&network) //编码
	decoder := gob.NewDecoder(&network) //解码
	// send data to encoder
	err := encoder.Encode(Stu{1, "zhangsan"})
	if err != nil {
		log.Fatal("编码异常")
	}
	var p Pel
	err = decoder.Decode(&p)
	if err != nil {
		log.Fatal("解码异常")
	}
	fmt.Println("解码结果", p)

}
