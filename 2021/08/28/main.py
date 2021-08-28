class Solution:
    def jobScheduling(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
        jobs = sorted(zip(startTime, endTime, profit), key=lambda l: l[1])
        dp = [[0, 0]]
        for s, e, p in jobs:
            pos = bisect_left(dp, [s + 1, 0]) - 1
            if dp[pos][1] + p > dp[-1][1]:
                dp.append([e, dp[pos][1] + p])
        return dp[-1][-1]