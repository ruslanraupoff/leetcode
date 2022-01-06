package main

func carPooling(trips [][]int, capacity int) bool {
	l := [1001]int{}
	for _, t := range trips {
		num, start, end := t[0], t[1], t[2]
		l[start] += num
		l[end] -= num
	}

	p := 0
	for _, c := range l {
		p += c
		if p > capacity {
			return false
		}
	}
	return true
}
