package main

func singleNumber(nums []int) []int {
	var x int
	for _, n := range nums {
		x ^= n
	}

	l := x & -x

	var a, b int
	for _, n := range nums {
		if n&l == 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a, b}
}
