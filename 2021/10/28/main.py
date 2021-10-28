class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        l, n = [], len(nums)
        nums.sort()
        for i in range(n - 2):
            if nums[i] > 0:
                break
            if i > 0 and nums[i] == nums[i-1]:
                continue
            t = -nums[i]
            j, k = i + 1, n - 1
            while j < k:
                s = nums[j] + nums[k]
                if s < t:
                    j += 1
                elif s > t:
                    k -= 1
                else:
                    l.append( [nums[i], nums[j], nums[k]] )
                    while j < n-1 and nums[j] == nums[j+1]:
                        j += 1
                    while k > 0 and nums[k] == nums[k-1]:
                        k -= 1
                    j += 1
                    k -= 1

        return l