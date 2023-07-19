package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Adderss struct {
	City    string
	Country string
}
type People struct {
	Id   int
	Name string
	Adds []*Adderss
}

func en2file() {
	add1 := &Adderss{"beijing", "China"}
	add2 := &Adderss{"nanjing", "China"}
	people := People{1, "zhangsan", []*Adderss{add1, add2}}
	file, _ := os.OpenFile("people.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(people)
	if err != nil {
		log.Fatal("编码错误")
	}

}
