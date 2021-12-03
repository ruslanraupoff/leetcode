package main

func maxProduct(nums []int) int {
	cur, neg, mx := 1, 1, nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			cur, neg = nums[i]*cur, nums[i]*neg
		} else if nums[i] < 0 {
			cur, neg = nums[i]*neg, nums[i]*cur
		} else {
			cur, neg = 0, 1
		}

		if mx < cur {
			mx = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return mx
}
