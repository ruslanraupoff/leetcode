class Solution:
    def maxTurbulenceSize(self, arr: List[int]) -> int:
        g, l, r = 1, 1, 1
        
        for i in range(1, len(arr)):
            if arr[i] > arr[i-1]:
                g, l = l + 1, 1
                r = max(r, g)
            elif arr[i] < arr[i-1]:
                g, l = 1, g+1
                r = max(r, l)
            else:
                g, l = 1, 1
        
        return r