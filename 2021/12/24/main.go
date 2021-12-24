package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int

	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			temp[1] = max(temp[1], intervals[i][1])
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
