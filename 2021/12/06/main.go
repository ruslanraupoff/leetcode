package main

func minCostToMoveChips(position []int) int {
	o, e := 0, 0
	for _, p := range position {
		if p%2 == 1 {
			e += 1
		} else {
			o += 1
		}
	}
	if o > e {
		return e
	}
	return o
}
