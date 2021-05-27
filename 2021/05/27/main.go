package main

import "strings"

func maxProduct(words []string) int {
	r := 0
	for i := 0; i < len(words)-1; i++ {
		a := words[i]
		la := len(a)
		for j := i + 1; j < len(words); j++ {
			b := words[j]
			lb := len(b)
			for k := 0; k < la; k++ {
				if strings.ContainsAny(b, string(a[k])) {
					lb = 0
					break
				}
			}
			if la*lb > r {
				r = la * lb
			}
		}
	}
	return r
}
