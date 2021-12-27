package main

func findComplement(num int) int {
	r, t := 0, num
	for t > 0 {
		t >>= 1
		r <<= 1
		r++
	}

	return r ^ num
}
