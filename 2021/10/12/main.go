package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)/2
		g := guess(m)
		if g == 0 {
			return m
		}
		if g == 1 {
			l = m + 1
		}
		if g == -1 {
			r = m - 1
		}
	}
	return -1
}
