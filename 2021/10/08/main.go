package main

import "strings"

type Trie struct {
	words []string
}

func Constructor() Trie {
	return Trie{words: make([]string, 0)}
}

func (this *Trie) Insert(word string) {
	this.words = append(this.words, word)
}

func (this *Trie) Search(word string) bool {
	for _, w := range this.words {
		if w == word {
			return true
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, w := range this.words {
		if strings.HasPrefix(w, prefix) {
			return true
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
