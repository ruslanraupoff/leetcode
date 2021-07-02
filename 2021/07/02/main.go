package main

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k
	for l < r {
		m := l + int((r-l)/2)
		if x-arr[m] <= arr[m+k]-x {
			r = m
		} else {
			l = m + 1
		}
	}
	return arr[l : l+k]
}
