package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	n := len(arr)

	minDiff := math.MaxInt64
	for i := 1; i < n; i++ {
		minDiff = min(minDiff, arr[i]-arr[i-1])
	}

	res := make([][]int, 0, n)

	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == minDiff {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
