package main

import (
	"math"
)

func arrangeCoins(n int) int {
	return int(math.Floor((-1.0 + math.Sqrt(float64(1+8*n))) * 0.5))
}
