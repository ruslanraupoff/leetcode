package main

func reverseOnlyLetters(s string) string {
	bs := []byte(s)

	l, r := 0, len(bs)-1
	for l < r {
		for l < r && !isLetter(bs[l]) {
			l++
		}
		for l < r && !isLetter(bs[r]) {
			r--
		}
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}

	return string(bs)
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}
