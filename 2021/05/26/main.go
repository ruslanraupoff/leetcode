package main

func minPartitions(n string) int {
	for i := 9; i > 0; i-- {
		for _, c := range n {
			if i+48 == int(c) {
				return i
			}
		}
	}
	return 0
}
