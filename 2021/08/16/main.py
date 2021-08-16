class NumArray:

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.sums = sum(nums)
        self.length = len(nums)
        

    def sumRange(self, left: int, right: int) -> int:
        a = right - left
        d = self.length - right + left - 1
        if a > d:
            return self.sums - sum(self.nums[:left]) - sum(self.nums[right+1:])
        return sum(self.nums[left:right+1]) 
        


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)