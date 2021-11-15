class Solution:
    def largestDivisibleSubset(self, nums: List[int]) -> List[int]:
        if not nums: return []
        dp = [0]*len(nums)
        nums.sort()
        idx, size= 0, 0
        for i in range(len(nums)):
            k, c= -1, 0
            for j in range(i-1,-1,-1):
                if nums[i]%nums[j] == 0 and len(dp[j])>c:
                    c = len(dp[j])
                    k = j 
            if k!=-1:
                dp[i] = dp[k]+[nums[i]]
            else:
                dp[i] = [nums[i]]

            if len(dp[i])>size:
                idx = i
                size = len(dp[i])
        return dp[idx]