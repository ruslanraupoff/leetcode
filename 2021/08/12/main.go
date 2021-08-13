package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	d := map[string][]string{}
	sort.Strings(strs)
	for _, w := range strs {
		key := sortWord(w)
		d[key] = append(d[key], w)
	}
	r := [][]string{}
	for _, g := range d {
		r = append(r, g)
	}
	return r
}

func sortWord(word string) string {
	byt := []byte(word)
	sort.Slice(byt, func(i, j int) bool {
		return byt[i] < byt[j]
	})
	return string(byt)
}
