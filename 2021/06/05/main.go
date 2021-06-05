package main

import (
	"container/heap"
	"sort"
)

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return efficiency[idxs[i]] > efficiency[idxs[j]]
	})
	ph := IntHeap{}
	heap.Init(&ph)
	res, sSum := 0, 0
	for _, index := range idxs {
		if ph.Len() == k {
			sSum -= heap.Pop(&ph).(int)
		}
		sSum += speed[index]
		heap.Push(&ph, speed[index])
		x := sSum * efficiency[index]
		if res < x {
			res = x
		}
	}
	return int(res % (1e9 + 7))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
