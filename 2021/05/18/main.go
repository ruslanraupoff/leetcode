package main

import (
	"strings"
)

func findDuplicate(paths []string) [][]string {
	d := make(map[string][]string, len(paths)*2)
	for _, p := range paths {
		ps := strings.Split(p, " ")
		for _, f := range ps[1:] {
			fs := strings.Split(f, "(")
			n, c := fs[0], fs[1][:len(fs[1])-1]
			d[c] = append(d[c], ps[0]+"/"+n)
		}
	}
	res := make([][]string, 0, len(d))
	for _, p := range d {
		if len(p) > 1 {
			res = append(res, p)
		}
	}
	return res
}
