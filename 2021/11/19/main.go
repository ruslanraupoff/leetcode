package main

func hammingDistance(x int, y int) int {
	d := 0
	v := x ^ y
	for v > 0 {
		v &= v - 1
		d += 1
	}
	return d
}
