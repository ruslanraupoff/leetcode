package main

func pushDominoes(dominoes string) string {
	r, i, ll, n := "", 0, 0, len(dominoes)
	for i < n {
		if dominoes[i] == '.' {
			i += 1
		} else if dominoes[i] == 'L' {
			for j := 0; j < i+1-ll; j++ {
				r += "L"
			}
			i += 1
			ll = i
		} else {
			for j := 0; j < i-ll; j++ {
				r += "."
			}
			lr := i
			i += 1
			for i < n {
				if dominoes[i] == '.' {
					i += 1
				} else if dominoes[i] == 'R' {
					for j := 0; j < i-lr; j++ {
						r += "R"
					}
					lr = i
					i += 1
				} else {
					t := (i - lr + 1)
					if t%2 == 1 {
						sr, sl := "", ""
						for j := 0; j < t/2; j++ {
							sr += "R"
							sl += "L"
						}
						r += sr + "." + sl
					} else {
						sr, sl := "", ""
						for j := 0; j < t/2; j++ {
							sr += "R"
							sl += "L"
						}
						r += sr + sl
					}
					i += 1
					lr = i
					break
				}
			}
			for j := 0; j < i-lr; j++ {
				r += "R"
			}
			ll = i
		}
	}
	for j := 0; j < i-ll; j++ {
		r += "."
	}
	return r
}
