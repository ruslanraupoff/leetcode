package main

func canReach(arr []int, start int) bool {

	if start < 0 || start > len(arr)-1 || arr[start] == -1 {
		return false
	}

	delta := arr[start]
	arr[start] = -1

	return (delta == 0) || canReach(arr, start+delta) || canReach(arr, start-delta)
}
