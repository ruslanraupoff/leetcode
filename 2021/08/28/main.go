package main

import (
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	sort.Sort(&jobs{startTime, endTime, profit})

	n := len(startTime)
	endInStart := make([]int, n)
	for i, v := range endTime {
		if v > startTime[n-1] {
			endInStart[i] = n
			continue
		}
		l, r := 0, n-1
		for l < r {
			mid := (l + r) / 2
			if v > startTime[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		endInStart[i] = l
	}
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		if profit[i]+dp[endInStart[i]] > dp[i] {
			dp[i] = profit[i] + dp[endInStart[i]]
		}
	}
	return dp[0]
}

type jobs struct {
	start, end, profit []int
}

func (job *jobs) Len() int {
	return len(job.start)
}

func (job *jobs) Swap(i, j int) {
	job.start[i], job.start[j] = job.start[j], job.start[i]
	job.end[i], job.end[j] = job.end[j], job.end[i]
	job.profit[i], job.profit[j] = job.profit[j], job.profit[i]
}

func (job *jobs) Less(i, j int) bool {
	return job.start[i] < job.start[j]
}
