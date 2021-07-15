package main

import "strings"

func customSortString(order string, str string) string {
	res := ""
	for i := 0; i < len(order); i++ {
		o := order[i : i+1]
		count := strings.Count(str, o)
		res += strings.Repeat(o, count)
		str = strings.Replace(str, o, "", -1)
	}

	return res + str
}
