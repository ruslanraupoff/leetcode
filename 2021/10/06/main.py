class Solution:
    def findDuplicates(self, nums: List[int]) -> List[int]:
        l, r = {}, []
        for n in nums:
            if n in l:
                r.append(n)
            else:
                l[n] = 1
        return r