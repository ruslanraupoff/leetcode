package main

func findDisappearedNumbers(nums []int) []int {
	l := make(map[int]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		l[i] = 0
	}
	for _, n := range nums {
		l[n] = 1
	}
	r := []int{}
	for i := 1; i <= len(nums); i++ {
		if l[i] == 0 {
			r = append(r, i)
		}
	}
	return r
}
