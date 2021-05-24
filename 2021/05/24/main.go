package main

func toLowerCase(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c = c + 'a' - 'A'
		}
		r += string(c)
	}
	return r
}
