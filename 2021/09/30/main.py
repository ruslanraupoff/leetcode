class Solution:
    def canPartitionKSubsets(self, nums: List[int], k: int) -> bool:
        s, n, mx = 0, len(nums), nums[0]
        for v in nums:
            s += v
            mx = max(mx, v)

        if s%k != 0 or s < mx*k:
            return False

        nums.sort(reverse=True)

        u = [False] * n

        def dfs(u, k, j, s, t):
            if k == 1:
                return True

            if s == t:
                return dfs(u, k-1, 0, 0, t)

            for i in range(j, n):
                if not u[i] and s+nums[i] <= t:
                    u[i] = True
                    if dfs(u, k, i+1, s+nums[i], t):
                        return True
                    u[i] = False

            return False

        return dfs(u, k, 0, 0, s//k)
