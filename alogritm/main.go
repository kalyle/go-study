package main

import (
	"fmt"
)

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple")) // 返回 True
	fmt.Println(trie.Search("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("applw"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
}
