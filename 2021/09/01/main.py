class Solution:
    def arrayNesting(self, nums: List[int]) -> int:
        r = 0
        for i in range(len(nums)):
            if nums[i] != -1:
                a = nums[i]
                c = 0
                while nums[a] != -1:
                    t = a
                    a = nums[a]
                    c += 1
                    nums[t] = -1
                r = max(r, c)
        return r
