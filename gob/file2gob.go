package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func de2file() {
	file, _ := os.Open("people.gob")
	defer file.Close()
	dec := gob.NewDecoder(file)
	var p People
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("解码错误")
	}
	fmt.Println("解码内容：", p, p.Adds[0])
}
