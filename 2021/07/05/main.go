package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if n*m != r*c {
		return mat
	}
	cnt, res := 0, make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = mat[cnt/m][cnt%m]
			cnt++
		}
	}
	return res
}
