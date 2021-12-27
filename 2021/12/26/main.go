package main

import "sort"

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i int, j int) bool {
		return dist(points[i]) < dist(points[j])
	})
	return points[:k]
}

func dist(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}
