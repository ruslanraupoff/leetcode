package main

type MapSum struct {
	mp map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{mp: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.mp[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	s, l := 0, len(prefix)
	for k, v := range this.mp {
		if len(k) >= len(prefix) && k[:l] == prefix {
			s += v
		}
	}
	return s
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
