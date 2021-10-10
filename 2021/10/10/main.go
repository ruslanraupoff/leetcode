package main

func rangeBitwiseAnd(left int, right int) int {
	r := 0
	for left >= 1 && right >= 1 {
		if left == right {
			return left << r
		}
		left >>= 1
		right >>= 1
		r += 1
	}

	return 0
}
