class Solution:
    def minPartitions(self, n: str) -> int:
        for x in range(9, 0, -1):
            if str(x) in n:
                return x