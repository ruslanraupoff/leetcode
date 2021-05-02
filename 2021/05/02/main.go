package main

import (
	"fmt"
	"sort"
)

type SortedList []int

func (list *SortedList) Insert(x int) {
	i := sort.SearchInts(*list, x)
	*list = append(*list, -1)
	copy((*list)[i+1:], (*list)[i:])
	(*list)[i] = x
}

func scheduleCourse(courses [][]int) int {
	qh := make(SortedList, 0)
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] })
	start := 0
	for i := 0; i < len(courses); i++ {
		dura, end := courses[i][0], courses[i][1]
		start += dura
		qh.Insert(dura)
		if start > end {
			start -= qh[len(qh)-1]
			qh = qh[:len(qh)-1]
		}
	}
	return len(qh)
}

func main() {
	var courses = [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
	var r = scheduleCourse(courses)
	fmt.Println(r)
}
