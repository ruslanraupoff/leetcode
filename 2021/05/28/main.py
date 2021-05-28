class Solution:
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        l, s, r, i = set(), 0, 0, 0
        for n in nums:
            while n in l:
                s -= nums[i]
                l.remove(nums[i])
                i += 1
            l.add(n)
            s += n
            r = max(r, s)
        return r