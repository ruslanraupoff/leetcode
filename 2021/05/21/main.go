package main

func findAndReplacePattern(words []string, pattern string) []string {
	var r []string
	nor := normalize(pattern)
	for _, w := range words {
		if nor == normalize(w) {
			r = append(r, w)
		}
	}
	return r
}

func normalize(str string) string {
	m := make(map[rune]byte, len(str))
	for _, c := range str {
		if _, ok := m[c]; !ok {
			m[c] = byte(len(m)) + 'a'
		}
	}
	res := make([]byte, len(str))
	for i, c := range str {
		res[i] = m[c]
	}
	return string(res)
}
