package main

const mod = 1e9 + 7

func nthMagicalNumber(n int, a int, b int) int {
	if a > b {
		a, b = b, a
	}

	m := lcm(a, b)
	l, r := a*n/2, b*n

	for {
		mid := (l + r) / 2
		count := magical_of(mid, a, b, m)
		if count < n {
			l = mid + 1
		} else if count > n {
			r = mid - 1
		} else {
			res := mid - min(mid%a, mid%b)
			return res % mod
		}
	}
}

func magical_of(num, a, b, m int) int {
	return num/a + num/b - num/m
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
