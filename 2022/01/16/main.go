package main

func maxDistToClosest(seats []int) int {
	d, e := 0, 0
	for i := 0; i < len(seats); i++ {
		if e == i {
			d = e
		} else {
			d = max(d, (e+e%2)/2)
		}
		if seats[i] == 1 {
			e = 0
		} else {
			e++
		}
	}

	return max(d, e)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
