package main

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, w)
	horizontalCuts = append(horizontalCuts, 0)
	verticalCuts = append(verticalCuts, 0)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	mxH, mxW, d := 0, 0, 0

	for i := 1; i < len(horizontalCuts); i++ {
		d = horizontalCuts[i] - horizontalCuts[i-1]
		if d > mxH {
			mxH = d
		}
	}
	for i := 1; i < len(verticalCuts); i++ {
		d = verticalCuts[i] - verticalCuts[i-1]
		if d > mxW {
			mxW = d
		}
	}

	return mxH * mxW % 1000000007
}
