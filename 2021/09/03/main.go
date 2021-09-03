package main

import "sort"

func outerTrees(trees [][]int) [][]int {
	if len(trees) <= 1 {
		return trees
	}

	bm := bottomLeft(trees)
	sort.Slice(trees, func(i int, j int) bool {
		diff := orientation(bm, trees[i], trees[j])
		if diff == 0 {
			return distance(bm, trees[i]) < distance(bm, trees[j])
		} else {
			return diff < 0
		}
	})

	i := len(trees) - 1
	for i >= 0 && orientation(bm, trees[len(trees)-1], trees[i]) == 0 {
		i--
	}

	j, k := i+1, len(trees)-1
	for j < k {
		trees[j], trees[k] = trees[k], trees[j]
		j++
		k--
	}

	stack := [][]int{trees[0], trees[1]}

	for p := 2; p < len(trees); p++ {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for orientation(stack[len(stack)-1], top, trees[p]) > 0 {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, top)
		stack = append(stack, trees[p])
	}
	return stack
}

func bottomLeft(trees [][]int) []int {
	res := trees[0]
	for _, v := range trees {
		if v[1] < res[1] {
			res = v
		}
	}
	return res
}

func orientation(p []int, q []int, r []int) int {
	return (q[1]-p[1])*(r[0]-q[0]) - (q[0]-p[0])*(r[1]-q[1])
}

func distance(p []int, q []int) int {
	return (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
}
