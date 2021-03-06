package main

func orderOfLargestPlusSign(n int, mines [][]int) int {
	k := n/2 + n%2

	if len(mines) == 0 {
		return k
	}

	isMined := [500][500]bool{}
	for _, m := range mines {
		isMined[m[0]][m[1]] = true
	}

	isGoodEdge := func(xBegin, xEnd, yBegin, yEnd int) bool {
		for i := xBegin; i <= xEnd; i++ {
			for j := yBegin; j <= yEnd; j++ {
				if isMined[i][j] {
					return false
				}
			}
		}
		return true
	}

	for ; k > 0; k-- {
		for i := k - 1; i <= n-k; i++ {
			for j := k - 1; j <= n-k; j++ {
				if !isMined[i][j] &&
					isGoodEdge(i-k+1, i-1, j, j) &&
					isGoodEdge(i+1, i+k-1, j, j) &&
					isGoodEdge(i, i, j-k+1, j-1) &&
					isGoodEdge(i, i, j+1, j+k-1) {
					return k
				}
			}
		}
	}

	return 0
}
