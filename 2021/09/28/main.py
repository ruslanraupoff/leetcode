class Solution:
    def sortArrayByParityII(self, nums: List[int]) -> List[int]:
        i, j, r = 0, 1, [0] * len(nums)
        for n in nums:
            if n % 2 == 0:
                r[i] = n
                i += 2
            else:
                r[j] = n
                j += 2
        return r