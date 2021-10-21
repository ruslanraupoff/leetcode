package main

func reverseWords(s string) string {
	buf := make([]byte, 0, len(s)/2)
	i := len(s) - 1
	count := 0
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		end := i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		if end > i {
			buf = append(buf, s[i+1:end+1]...)
			buf = append(buf, ' ')
			count++
		}
	}
	if count > 0 {
		buf = buf[:len(buf)-1]
	}
	return string(buf)
}
