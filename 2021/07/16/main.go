package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, l := j+1, len(nums)-1; k < l; {
				s := nums[i] + nums[j] + nums[k] + nums[l]
				if s == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					l--
					for k < l && nums[l] == nums[l+1] {
						l--
					}
					k++
					for k < l && nums[k] == nums[k-1] {
						k++
					}
				} else if s > target {
					l--
				} else {
					k++
				}
			}
		}
	}

	return res
}
