package main

import "strings"

func numMatchingSubseq(s string, words []string) int {
	r, hs, uhs := 0, make(map[string]bool), make(map[string]bool)

	for _, w := range words {
		_, existh := hs[w]
		if existh {
			r++
			continue
		}
		_, existu := uhs[w]
		if existu {
			continue
		}
		if is_sub(s, w) {
			hs[w] = true
			r++
		} else {
			uhs[w] = true
		}
	}
	return r
}

func is_sub(s string, word string) bool {
	idx := 0
	for _, w := range word {
		fidx := strings.IndexRune(s[idx:], w)
		if fidx == -1 {
			return false
		}
		idx += fidx + 1
	}
	return true
}
