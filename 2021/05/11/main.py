class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        t = sum(cardPoints[:k])
        r = t
        for i in range(k):
            t = t - cardPoints[k-i-1] + cardPoints[-i-1]
            r = max(r, t)
        return r
