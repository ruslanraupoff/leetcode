package main

func minFlipsMonoIncr(s string) int {
	o, f := 0, 0
	for _, c := range s {
		if c == '0' {
			f = min(f+1, o)
		} else {
			o += 1
		}
	}
	return f
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
