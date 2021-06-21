package main

func generate(numRows int) [][]int {
	r := [][]int{}
	for i := 1; i <= numRows; i++ {
		l := make([]int, i)
		for j := range l {
			l[j] = 1
		}
		for j := 1; j+1 < i; j++ {
			l[j] = r[i-2][j-1] + r[i-2][j]
		}
		r = append(r, l)
	}
	return r
}
