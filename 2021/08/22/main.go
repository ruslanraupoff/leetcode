package main

import "sort"

func rectangleArea(rectangles [][]int) int {
	size := len(rectangles)

	xs := make([]int, 0, size*2)
	m := make(map[int]bool, size*2)
	for _, r := range rectangles {
		m[r[0]] = true
		m[r[2]] = true
	}
	for k := range m {
		xs = append(xs, k)
	}
	sort.Ints(xs)
	idxs := make(map[int]int, 2*len(xs))
	for idx, x := range xs {
		idxs[x] = idx
	}

	rect := make([][]int, 0, 2*size)
	for _, v := range rectangles {
		x1, y1, x2, y2 := v[0], v[1], v[2], v[3]
		rect = append(rect,
			[]int{y1, x1, x2, 1},
			[]int{y2, x1, x2, -1},
		)
	}
	sort.Slice(rect, func(i int, j int) bool {
		return rect[i][0] < rect[j][0]
	})

	c := make([]int, len(xs))

	cy, s, r := 0, 0, 0

	for _, l := range rect {
		y, x1, x2, sig := l[0], l[1], l[2], l[3]
		r += (y - cy) * s
		cy = y
		s = 0
		for i := idxs[x1]; i < idxs[x2]; i++ {
			c[i] += sig
		}
		for i := 0; i+1 < len(c); i++ {
			if c[i] > 0 {
				s += xs[i+1] - xs[i]
			}
		}
	}

	return r % (1e9 + 7)
}
