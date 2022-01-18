package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	i, m := 0, len(flowerbed)
	for n > 0 && i < m {
		if flowerbed[i] == 1 {
			i += 2
		} else if i < m-1 && flowerbed[i+1] == 1 {
			i += 3
		} else {
			n -= 1
			i += 2
		}
	}

	return n == 0
}
