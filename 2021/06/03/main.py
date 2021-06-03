class Solution:
    def maxArea(self, h: int, w: int, horizontalCuts: List[int], verticalCuts: List[int]) -> int:
        horizontalCuts.append(0)
        horizontalCuts.append(h)
        verticalCuts.append(0)
        verticalCuts.append(w)
        horizontalCuts.sort()
        verticalCuts.sort()
        mxH, mxW = 0, 0
        
        for i in range(1, len(horizontalCuts)):
            mxH = max(mxH, horizontalCuts[i] - horizontalCuts[i-1])
        for i in range(1, len(verticalCuts)):
            mxW = max(mxW, verticalCuts[i] - verticalCuts[i-1])

        return mxH * mxW % (10**9 + 7)
