package main

import "regexp"

func isNumber(s string) bool {
	valid, _ := regexp.MatchString(`^[\+\-]?(\d+|\d+\.\d*|\.\d+)([eE][\+\-]?\d+)?$`, s)
	return valid
}
