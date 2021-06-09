class Solution:
    def maxResult(self, nums: List[int], k: int) -> int:
        dp = -nums[-1]
        pq = [(dp, len(nums) - 1)]
        for i in range(len(nums) - 2, -1, -1):
            while pq[0][1] > i + k:
                heapq.heappop(pq)
            dp = pq[0][0] - nums[i]
            heapq.heappush(pq, (dp, i))
        return -dp
