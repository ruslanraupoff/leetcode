class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        dp = []
        for i in range(len(nums)):
            idx = binary_search(dp, nums[i])
            k = len(dp)
            if idx == k:
                arr = []
                if k != 0:
                    arr = [i for i in dp[-1]]
                arr.append(nums[i])
                dp.append(arr)
            elif dp[idx][-1] > nums[i]:
                dp[idx][-1] = nums[i]

        return len(dp[-1])

def binary_search(dp, target):
    l = 0
    r = len(dp) - 1
    while l <= r:
        mid = l + (r - l) // 2
        if dp[mid][-1] == target:
            return mid
        elif dp[mid][-1] < target:
            l = mid + 1
        else:
            r = mid - 1
    return l