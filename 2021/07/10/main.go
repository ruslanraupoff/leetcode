package main

func numDecodings(s string) int {
	l, r, m := 1, 1, 1000000007
	if s[0] == '*' {
		r = 9
	} else if s[0] == '0' {
		r = 0
	}
	for i := 1; i < len(s); i++ {
		t := r
		if s[i] == '*' {
			r = 9 * r % m
			if s[i-1] == '1' {
				r = (r + 9*l) % m
			} else if s[i-1] == '2' {
				r = (r + 6*l) % m
			} else if s[i-1] == '*' {
				r = (r + 15*l) % m
			}
		} else {
			if s[i] == '0' {
				r = 0
			}
			if s[i-1] == '1' {
				r = (r + l) % m
			} else if s[i-1] == '2' && s[i] <= '6' {
				r = (r + l) % m
			} else if s[i-1] == '*' {
				if s[i] <= '6' {
					r = (r + 2*l) % m
				} else {
					r = (r + l) % m
				}
			}
		}
		l = t
	}
	return r
}
