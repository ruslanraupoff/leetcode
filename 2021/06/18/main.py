class NumArray:

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.sums = sum(nums)
        self.length = len(nums)

    def update(self, index: int, val: int) -> None:
        self.sums = self.sums - self.nums[index] + val
        self.nums[index] = val

    def sumRange(self, left: int, right: int) -> int:
        a = right - left
        d = self.length - right + left - 1
        if a > d:
            return self.sums - sum(self.nums[:left]) - sum(self.nums[right+1:])
        return sum(self.nums[left:right+1]) 


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# obj.update(index,val)
# param_2 = obj.sumRange(left,right)