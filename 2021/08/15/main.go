package main

func minWindow(s string, t string) string {
	a := make([]int, 58)
	for i := 0; i < 58; i++ {
		a[i] = 0
	}
	for _, c := range t {
		k := c - 'A'
		a[k] = a[k] + 1
	}
	l, r, mx := 0, 0, 1000000
	i, j, n := 0, -1, len(t)
	for l < len(s) {
		k := s[r] - 'A'
		a[k] -= 1
		if a[k] >= 0 {
			n -= 1
		}
		for n == 0 {
			m := r - l + 1
			if m < mx {
				i, j, mx = l, r, m
			}
			k = s[l] - 'A'
			a[k] += 1
			if a[k] > 0 {
				n += 1
			}
			l += 1
		}
		r += 1
	}

	return s[i : j+1]
}
