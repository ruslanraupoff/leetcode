package main

type tree struct {
	next   [26]*tree
	isWord bool
}

type StreamChecker struct {
	tree   *tree
	stream []int
}

func Constructor(words []string) StreamChecker {
	m := 0
	t := &tree{}
	for _, w := range words {
		m = max(m, len(w))
		t.insert(w)
	}
	return StreamChecker{
		tree:   t,
		stream: make([]int, 0, 1024),
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.stream = append(this.stream, int(letter-'a'))
	n, t := len(this.stream), this.tree
	for i := n - 1; i >= 0; i-- {
		index := this.stream[i]
		if t.next[index] == nil {
			return false
		}
		t = t.next[index]
		if t.isWord {
			return true
		}
	}
	return false
}

func (t *tree) insert(word string) {
	n := len(word)
	// reversely insert
	for i := n - 1; i >= 0; i-- {
		index := int(word[i] - 'a')
		if t.next[index] == nil {
			t.next[index] = &tree{}
		}
		t = t.next[index]
	}
	t.isWord = true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
