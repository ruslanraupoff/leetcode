package main

func ambiguousCoordinates(s string) []string {
	var r []string
	for i := 2; i+1 < len(s); i++ {
		for _, x := range maked(s, 1, i) {
			for _, y := range maked(s, i, len(s)-1) {
				r = append(r, "("+x+", "+y+")")
			}
		}
	}
	return r
}

func maked(s string, i int, n int) []string {
	var r []string
	for j := 1; j <= n-i; j++ {
		left, right := s[i:i+j], s[i+j:n]
		if (left[0] != '0' || left == "0") && (right == "" || right[len(right)-1] != '0') {
			if right != "" {
				r = append(r, left+"."+right)
			} else {
				r = append(r, left+right)
			}
		}
	}
	return r
}
