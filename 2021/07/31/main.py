class Solution:
    def trap(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1
        lmx, rmx, t = 0, 0, 0
        while l < r:
            lmx = max(lmx, height[l])
            rmx = max(rmx, height[r])

            if lmx < rmx:
                t += max(lmx - height[l+1], 0)
                l += 1
            else:
                t += max(rmx - height[r-1], 0)
                r -= 1
        return t