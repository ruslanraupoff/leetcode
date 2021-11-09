package main

func findNumOfValidWords(words []string, puzzles []string) []int {
	mp := make(map[int]int, len(words))
	for _, w := range words {
		mp[mask(w)]++
	}

	res := make([]int, len(puzzles))

	for i, p := range puzzles {
		fb := 1 << uint(p[0]-'a')
		ob := mask(p)
		sb := ob
		for sb != 0 {
			if sb&fb != 0 {
				res[i] += mp[sb]
			}
			sb = (sb - 1) & ob
		}
	}

	return res
}

func mask(w string) int {
	res := 0
	for _, l := range w {
		res |= 1 << uint(l-'a')
	}
	return res
}
