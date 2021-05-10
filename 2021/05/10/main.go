package main

func countPrimes(n int) int {
	primes := make([]bool, n+1)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if !primes[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			primes[j] = false
		}
	}
	r := 0
	for i := 2; i < n; i++ {
		if primes[i] {
			r++
		}
	}
	return r
}
