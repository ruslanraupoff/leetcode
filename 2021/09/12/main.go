package main

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	flag := true
	steps := make([]int, n)
	for i := 0; i < n; i++ {
		steps[i] = -1
	}
	steps[0] = maxMoves
	for flag {
		flag = false
		for _, edge := range edges {
			if steps[edge[0]] > 0 {
				na := steps[edge[0]] - edge[2] - 1
				if na > steps[edge[1]] {
					steps[edge[1]] = na
					flag = true
				}
			}
			if steps[edge[1]] > 0 {
				na := steps[edge[1]] - edge[2] - 1
				if na > steps[edge[0]] {
					steps[edge[0]] = na
					flag = true
				}
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if steps[i] != -1 {
			res += 1
		}
	}
	for _, edge := range edges {
		tmp := 0
		if steps[edge[0]] > 0 {
			tmp += steps[edge[0]]
		}
		if steps[edge[1]] > 0 {
			tmp += steps[edge[1]]
		}

		if tmp > edge[2] {
			res += edge[2]
		} else {
			res += tmp
		}
	}

	return res
}
