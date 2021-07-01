package main

func grayCode(n int) []int {
	return gen([]int{0}, n, 1)
}

func gen(r []int, n, b int) []int {
	if n == 0 {
		return r
	}
	l := len(r)
	t := make([]int, l)
	for i := range r {
		t[l-i-1] = r[i] + b
	}

	return gen(append(r, t...), n-1, b*2)
}
