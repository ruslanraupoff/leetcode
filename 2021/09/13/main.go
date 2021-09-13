package main

func maxNumberOfBalloons(text string) int {
	p := map[rune]int{'a': 0, 'b': 0, 'l': 0, 'n': 0, 'o': 0}
	for _, t := range text {
		if _, ok := p[t]; !ok {
			p[t] += 1
		}
	}
	mn := min(p['b'], min(p['a'], p['n']))
	mn = min(mn, min(p['l']/2, p['o']/2))
	return mn
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
