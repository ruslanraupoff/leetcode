package main

type WordFilter struct {
	mp map[string]int
}

func Constructor(words []string) WordFilter {
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		w := words[i]
		for j := 1; j <= len(w); j++ {
			for k := 1; k <= len(w); k++ {
				m[w[:j]+"#"+w[len(w)-k:]] = i
			}
		}
	}
	return WordFilter{mp: m}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	p := prefix + "#" + suffix
	if index, ok := this.mp[p]; ok {
		return index
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
