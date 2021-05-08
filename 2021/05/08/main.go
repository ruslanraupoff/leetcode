package main

import "strconv"

func superpalindromesInRange(left string, right string) int {
	r := 0
	var l []int
	for i := 1; i < 10; i++ {
		l = append(l, i)
	}

	for i := 1; i < 10000; i++ {
		str := strconv.Itoa(i)
		rev := reverse(str)
		v, _ := strconv.Atoi(str + rev)
		l = append(l, v)
		for j := 0; j < 10; j++ {
			mid := strconv.Itoa(j)
			v, _ := strconv.Atoi(str + mid + rev)
			l = append(l, v)
		}
	}

	for _, x := range l {
		v := x * x
		start, _ := strconv.Atoi(left)
		end, _ := strconv.Atoi(right)
		if start <= v && v <= end {
			if ispalindrome(strconv.Itoa(v)) {
				r++
			}
		}
	}

	return r
}

func ispalindrome(s string) bool {
	return s == reverse(s)
}

func reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}

func main() {

}
