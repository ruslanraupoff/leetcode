package main

func beautifulArray(n int) []int {
	r := []int{1}
	for len(r) < n {
		tmp := make([]int, 0, len(r)*2)
		for _, v := range r {
			if v*2-1 <= n {
				tmp = append(tmp, v*2-1)
			}
		}
		for _, v := range r {
			if v*2 <= n {
				tmp = append(tmp, v*2)
			}
		}
		r = tmp
	}
	return r
}
