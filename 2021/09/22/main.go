package main

import "math"

func maxLength(arr []string) int {
	return explore(arr, 0, 0)
}

func explore(a []string, j int, u int) int {
	if j == len(a) {
		return 0
	}

	cu := 0
	s := a[j]
	v := true
	for i := 0; i < len(s); i++ {
		ex := cu & (1 << (s[i] - 'a'))
		if ex == 0 {
			cu |= 1 << (s[i] - 'a')
		} else {
			v = false
			break
		}
	}

	r := explore(a, j+1, u)
	if v && (u&cu) == 0 {
		r = int(math.Max(float64(r),
			float64(len(s)+explore(a, j+1, u|cu))))
	}

	return r
}
