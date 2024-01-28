package main

import (
	"fmt"
)

const size = 26

type Node struct {
	children [size]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func (t *Trie) insert(s string) {
	n := len(s)
	currentNode := t.root
	for i := 0; i < n; i++ {
		charIdx := s[i] - 'a'
		if currentNode.children[charIdx] == nil {
			currentNode.children[charIdx] = &Node{}
		}
		currentNode = currentNode.children[charIdx]
	}
	currentNode.isEnd = true
}
func (t *Trie) search(s string) bool {
	n := len(s)
	currentNode := t.root
	for i := 0; i < n; i++ {
		charIdx := s[i] - 'a'
		if currentNode.children[charIdx] == nil {
			return false
		}
		currentNode = currentNode.children[charIdx]
	}

	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	t := &Trie{root: &Node{}}
	toAdd := []string{
		"aravindakshan",
		"aravind",
		"arav",
		"ara",
		"ar",
	}
	for _, v := range toAdd {
		t.insert(v)
	}
	b := t.search("akhil")
	fmt.Println(b)
	c := t.search("h")
	fmt.Println(c)

}
