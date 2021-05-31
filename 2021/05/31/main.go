package main

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	var a [][]string
	for i := 1; i <= len(searchWord); i++ {
		q := searchWord[:i]
		j := 0
		var st []string
		for j < len(products) {
			p := products[j]
			if strings.HasPrefix(p, q) {
				st = append(st, p)
			}
			j += 1
			if len(st) == 3 {
				break
			}
		}
		a = append(a, st)
	}
	return a
}
