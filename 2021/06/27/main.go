package main

func candy(ratings []int) int {
	a, n := 0, len(ratings)
	l, r := make([]int, n), make([]int, n)
	l[0], r[n-1] = 1, 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			l[i] = l[i-1] + 1
		} else {
			l[i] = 1
		}
		if ratings[n-i-1] > ratings[n-i] {
			r[n-i-1] = r[n-i] + 1
		} else {
			r[n-i-1] = 1
		}
	}
	for i := 0; i < n; i++ {
		if l[i] > r[i] {
			a += l[i]
		} else {
			a += r[i]
		}
	}
	return a
}
