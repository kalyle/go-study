package main

import "fmt"

// 前缀树
// type TrieNode struct {
// 	Count    uint64 //当前字符出现频率,以当前单词结尾的单词数量
// 	Prefix   string //以该处节点之前的字符串为前缀的单词数量
// 	NextNode *[26]TrieNode
// }

// func (trei *TrieNode) Insert(root *TrieNode, word string) {
// 	if root == nil || len(word) == 0 {
// 		return
// 	}
// 	list := []rune(word)
// 	for i := 0; i < len(list); i++ {
// 		// 分支是否存在
// 		if root.NextNode[list[i]-'a'] {
// 			root.NextNode[list[i]-'a'] = TrieNode{}
// 		}
// 	}
// }
// func Search(trie *TrieNode, word string) bool {
// 	return false
// }
type Trie struct {
	children map[string]*Trie
	end      bool
}

func Constructor() Trie {
	children := make(map[string]*Trie)
	return Trie{children: children, end: false}
}

func (this *Trie) Insert(word string) {
	current := this
	for _, val := range word {
		// if children == nil {
		// 	children = make(map[string]*Trie)
		// } 无用
		str := string(val)
		if _, ok := current.children[str]; !ok {
			// *children[str] = Constructor() error
			node := Constructor()
			current.children[str] = &node

		}
		current = current.children[str]
	}
	current.end = true

}

func (this *Trie) Search(word string) bool {
	current := this
	for _, val := range word {
		str := string(val)
		if _, ok := current.children[str]; !ok {
			return false

		}
		current = current.children[str]
	}
	return current.end
}

func (this *Trie) StartsWith(prefix string) bool {
	children := this.children
	for _, val := range prefix {
		str := string(val)
		if _, ok := children[str]; !ok {
			return false

		}
		children = children[str].children
	}
	return true
}

func (this *Trie) Show() {
	fmt.Println("")
}
