package main

import "sort"

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i int, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 0
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] <= end {
			continue
		}
		res++
		end = points[i][1]
	}

	res++

	return res
}
