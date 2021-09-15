package main

func maxTurbulenceSize(arr []int) int {
	g, l, r := 1, 1, 1
	for i := 1; i < len(arr); i++ {
		switch {
		case arr[i] > arr[i-1]:
			g, l = l+1, 1
			r = max(r, g)
		case arr[i] < arr[i-1]:
			g, l = 1, g+1
			r = max(r, l)
		default:
			g, l = 1, 1
		}
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
