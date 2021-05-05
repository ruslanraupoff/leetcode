package main

import "fmt"

func jump(nums []int) int {
	r, mx, mj := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if mx < i+nums[i] {
			mx = i + nums[i]
		}
		if i == mj {
			r += 1
			mj = mx
		}
	}
	return r
}

func main() {
	var res = jump([]int{2, 3, 1, 1, 4})
	fmt.Println(res)
}
