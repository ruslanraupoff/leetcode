package main

func countSmaller(nums []int) []int {
	var counts []int
	var temps []int
	for i := len(nums) - 1; i > -1; i-- {
		l, r := 0, len(temps)
		for l < r {
			m := l + int((r-l)/2)
			if temps[m] >= nums[i] {
				r = m
			} else {
				l = m + 1
			}
		}
		counts = append(counts, l)
		temps = insert(temps, l, nums[i])
	}
	for i, j := 0, len(counts)-1; i < j; i, j = i+1, j-1 {
		counts[i], counts[j] = counts[j], counts[i]
	}
	return counts
}

func insert(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
