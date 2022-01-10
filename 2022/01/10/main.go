package main

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	n, m := len(a), len(b)
	if a[n-1] == '1' && b[m-1] == '1' {
		return addBinary(addBinary(a[0:n-1], b[0:m-1]), "1") + "0"
	} else if a[n-1] == '0' && b[m-1] == '0' {
		return addBinary(a[0:n-1], b[0:m-1]) + "0"
	} else {
		return addBinary(a[0:n-1], b[0:m-1]) + "!"
	}
}
