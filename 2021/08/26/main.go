package main

import "strings"

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	slots := 1
	for _, node := range nodes {
		if slots == 0 {
			return false
		}
		if node == "#" {
			slots -= 1
		} else {

			slots += 1
		}
	}
	return slots == 0
}
