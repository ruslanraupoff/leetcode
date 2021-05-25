package main

import "strconv"

func evalRPN(tokens []string) int {
	var l []int
	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" {
			d := len(l) - 1
			b, a := l[d], l[d-1]
			l = l[:d-1]
			l = append(l, calc(a, b, t))
		} else {
			v, _ := strconv.Atoi(t)
			l = append(l, v)
		}
	}
	return l[0]
}

func calc(a, b int, o string) int {
	switch o {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
