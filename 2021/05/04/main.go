package main

import "fmt"

func checkPossibility(nums []int) bool {
	var cnt = 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			cnt++
			if cnt > 1 || i > 0 && nums[i-1] > nums[i+1] && i+2 < len(nums) && nums[i+2] < nums[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	var res = checkPossibility([]int{4, 2, 1})
	fmt.Println(res)
}
