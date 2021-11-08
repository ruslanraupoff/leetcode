package main

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if n == 3 {
		return 5
	}

	r := 0
	for i := 1; i <= n/2; i++ {
		r += numTrees(i-1) * numTrees(n-i)
	}
	r *= 2
	if n%2 == 1 {
		t := numTrees(n / 2)
		r += t * t
	}

	return r
}
