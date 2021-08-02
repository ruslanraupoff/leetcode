package main

func twoSum(nums []int, target int) []int {
	hs := map[int]int{}
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		_, exists := hs[n]
		if exists {
			return []int{hs[n], i}
		}
		hs[target-n] = i
	}
	return []int{-1, -1}
}
