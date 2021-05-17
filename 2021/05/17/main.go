package main

import "sort"

func longestStrChain(words []string) int {
	mp := map[string]int{}
	ans := 0
	for _, w := range words {
		mp[w] = 1
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, w := range words {
		for i := 0; i < len(w); i++ {
			p := w[:i] + w[i+1:]
			_, exists := mp[p]
			if exists && mp[p]+1 > mp[w] {
				mp[w] = mp[p] + 1
			}
		}
		if mp[w] > ans {
			ans = mp[w]
		}
	}
	return ans
}
