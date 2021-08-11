package main

import "sort"

func canReorderDoubled(arr []int) bool {
	mp := make(map[int]int)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]*arr[i] < arr[j]*arr[j]
	})
	for _, v := range arr {
		mp[v]++
	}
	for _, n := range arr {
		if mp[n] <= 0 {
			continue
		}
		_, ok := mp[2*n]
		if ok && mp[2*n] > 0 {
			mp[n] -= 1
			mp[2*n] -= 1
		} else {
			return false
		}
	}
	return true
}
