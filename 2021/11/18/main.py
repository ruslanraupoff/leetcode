class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        l = {i: 0 for i in range(1, len(nums)+1)}
        for n in nums:
            l[n] = 1
        r = []
        for i in range(1, len(nums)+1):
            if l[i] == 0:
                r.append(i)
        return r