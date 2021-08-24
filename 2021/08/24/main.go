package main

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	a, b := parse(num1)
	c, d := parse(num2)

	return strconv.Itoa(a*c-b*d) + "+" + strconv.Itoa(a*d+b*c) + "i"
}

func parse(y string) (re, im int) {
	z := strings.Split(y, "+")
	re, _ = strconv.Atoi(z[0])
	im, _ = strconv.Atoi(z[1][:len(z[1])-1])
	return
}
