package main

import "sort"

func threeSum(nums []int) [][]int {
	l, n := [][]int{}, len(nums)
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k, t := i+1, n-1, -nums[i]

		for j < k {
			s := nums[j] + nums[k]
			if s < t {
				j += 1
			} else if s > t {
				k -= 1
			} else {
				l = append(l, []int{nums[i], nums[j], nums[k]})
				for j < n-1 && nums[j] == nums[j+1] {
					j += 1
				}
				for k > 0 && nums[k] == nums[k-1] {
					k -= 1
				}
				j += 1
				k -= 1
			}
		}
	}

	return l
}
