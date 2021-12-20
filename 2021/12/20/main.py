class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        md = 10**6+1
        arr.sort()
        for i in range(1, len(arr)):
            md = min(md, arr[i]-arr[i-1])
        r = []
        for i in range(1, len(arr)):
            if arr[i]-arr[i-1] == md:
                r.append([arr[i-1], arr[i]])
        return r