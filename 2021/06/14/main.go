package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	cnt, res := 0, 0
	for i := 0; i < len(boxTypes); i++ {
		for j := 0; j < boxTypes[i][0]; j++ {
			if cnt < truckSize {
				res += boxTypes[i][1]
				cnt += 1
			}
		}
	}
	return res
}
