package main

type NumArray struct {
	sums   int
	nums   []int
	length int
}

func Constructor(nums []int) NumArray {
	s, l := 0, len(nums)
	for i := 0; i < l; i++ {
		s += nums[i]
	}
	return NumArray{
		sums:   s,
		length: l,
		nums:   nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.sums = this.sums - this.nums[index] + val
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	s, a, d := 0, right-left, this.length-right+left-1
	if a > d {
		for i := 0; i < left; i++ {
			s += this.nums[i]
		}
		for i := right + 1; i < this.length; i++ {
			s += this.nums[i]
		}
		return this.sums - s
	}
	for i := left; i <= right; i++ {
		s += this.nums[i]
	}
	return s
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
