package main

func sortArrayByParityII(nums []int) []int {
	i, j, r := 0, 1, make([]int, len(nums))
	for _, n := range nums {
		if n%2 == 0 {
			r[i] = n
			i += 2
		} else {
			r[j] = n
			j += 2
		}
	}
	return r
}
