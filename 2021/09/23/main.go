package main

func breakPalindrome(palindrome string) string {
	i, n := 0, len(palindrome)
	if n == 1 {
		return ""
	}
	l := []byte(palindrome)
	for i < n {
		if l[i] != 'a' && i != n-i-1 {
			l[i] = 'a'
			break
		}
		i += 1
	}
	if i == n {
		l[i-1] = byte(int(l[i-1]) + 1 + 'a')
	}

	return string(l)
}
