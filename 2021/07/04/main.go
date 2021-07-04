package main

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1

	for j := 1; j < n; j++ {
		a, e, i, o, u = e%1000000007, (a+i)%1000000007, (a+e+o+u)%1000000007, (i+u)%1000000007, a%1000000007
	}

	return (a + e + i + o + u) % 1000000007
}
