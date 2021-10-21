package main

import "math/rand"

type RandomizedSet struct {
	a []int
	m map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		a: []int{},
		m: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.a = append(this.a, val)
	this.m[val] = len(this.a) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}

	this.a[this.m[val]] = this.a[len(this.a)-1]
	this.m[this.a[len(this.a)-1]] = this.m[val]
	this.a = this.a[:len(this.a)-1]
	delete(this.m, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.a[rand.Intn(len(this.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
