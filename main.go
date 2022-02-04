package main

import (
	"fmt"
)

func main() {
	in := []string{"ab", "ba", "aaab", "abab", "baa"}
	s := Constructor(in)
	b := []byte{'a', 'a', 'a'}
	for i := 0; i < len(b); i++ {
		fmt.Println(s.Query(b[i]))
	}
	fmt.Println(s)
}

type StreamChecker struct {
	in   []byte
	root *tree
}

type tree struct {
	c        byte
	w        bool
	children []*tree
}

func Constructor(words []string) StreamChecker {
	root := &tree{}
	for _, word := range words {
		root.insert(word)
	}

	return StreamChecker{in: make([]byte, 0), root: root}
}

func (root *tree) insert(word string) {
	t := root
	for i := len(word) - 1; i >= 0; i-- {
		if t.children == nil {
			t.children = make([]*tree, 0)
		}

		k := -1
		for j := 0; j < len(t.children); j++ {
			if t.children[j].c == word[i] {
				k = j
			}
		}

		if k == -1 {
			t.children = append(t.children, &tree{c: word[i]})
			k = len(t.children) - 1
		}

		t = t.children[k]
	}

	t.w = true
}

func (root *tree) search(in []byte) bool {
	t := root
	for i := len(in) - 1; i >= 0; i-- {
		k := -1
		for j := 0; j < len(t.children); j++ {
			if t.children[j].c == in[i] {
				k = j
			}
		}

		if k == -1 {
			return false
		}

		t = t.children[k]
		if t.w {
			return true
		}
	}

	return false
}

func (sc *StreamChecker) Query(letter byte) bool {
	sc.in = append(sc.in, letter)
	return sc.root.search(sc.in)
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
