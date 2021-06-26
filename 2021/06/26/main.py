class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        counts, temps = [], []
        for i in range(len(nums) - 1, -1, -1):
            l, r = 0, len(temps)
            while l < r:
                m = l + (r-l)//2
                if temps[m] >= nums[i]:
                    r = m
                else:
                    l = m + 1
            counts.append(l)
            temps.insert(l, nums[i])
        counts.reverse()
        return counts