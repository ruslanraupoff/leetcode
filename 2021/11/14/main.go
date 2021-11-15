package main

type CombinationIterator struct {
	combines []string
	index    int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combines := make([]string, 0)
	dfs(characters, 0, combinationLength, "", &combines)
	return CombinationIterator{
		combines: combines,
		index:    0,
	}
}

func dfs(characters string, start, combinationLength int, currStr string, res *[]string) {
	if len(currStr) == combinationLength {
		*res = append(*res, currStr)
		return
	}
	for i := start; i < len(characters); i++ {
		currStr = currStr + characters[i:i+1]
		dfs(characters, i+1, combinationLength, currStr, res)
		currStr = currStr[:len(currStr)-1]
	}
}

func (this *CombinationIterator) Next() string {
	res := this.combines[this.index]
	this.index++
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return this.index != len(this.combines)
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
