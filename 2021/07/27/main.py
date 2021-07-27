class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        d, ln = 10 ** 5, len(nums)
        nums.sort()
        for i in range(ln):
            l, r = i + 1, ln - 1
            while l < r:
                sum = nums[i] + nums[l] + nums[r]
                if abs(target - sum) < abs(d):
                    d = target - sum
                if sum < target:
                    l += 1
                else:
                    r -= 1
            if d == 0:
                break
        return target - d