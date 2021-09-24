package main

func tribonacci(n int) int {
	t := []int{0, 1, 1}
	for i := 0; i < n; i++ {
		t = append(t, t[i]+t[i+1]+t[i+2])
	}
	return t[n]
}
