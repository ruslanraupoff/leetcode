package main

func palindromePairs(words []string) [][]int {
	var res [][]int
	hm := make(map[string]int, 0)
	for i := 0; i < len(words); i++ {
		hm[words[i]] = i
	}
	for i := 0; i < len(words); i++ {
		word := words[i]
		n := len(word)
		for j := 0; j <= n; j++ {
			pre, suf := word[:j], word[j:]
			if ispalindrome(pre) {
				rstr := reverse(suf)
				_, exists := hm[rstr]
				if exists && rstr != word {
					res = append(res, []int{hm[rstr], i})
				}
			}
			if j != n && ispalindrome(suf) {
				rstr := reverse(pre)
				_, exists := hm[rstr]
				if exists && rstr != word {
					res = append(res, []int{i, hm[rstr]})
				}
			}
		}
	}
	return res
}

func ispalindrome(s string) bool {
	return s == reverse(s)
}

func reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}
