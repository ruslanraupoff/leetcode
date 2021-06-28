package main

func removeDuplicates(s string) string {
	l := len(s)
	t := []byte{}
	for i := 0; i < l; i++ {
		lt := len(t)
		if lt != 0 && s[i] == t[lt-1] {
			t = t[:lt-1]
		} else {
			t = append(t, s[i])
		}
	}
	return string(t)
}
