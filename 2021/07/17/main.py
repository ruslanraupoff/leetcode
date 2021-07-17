class Solution:
    def threeEqualParts(self, arr: List[int]) -> List[int]:
        n, s = len(arr), sum(arr)

        if s == 0:
            return [0, n - 1]

        if s % 3:
            return [-1, -1]

        k, j, t, i = s // 3, n - 1, 0, 0
        while j >= 0:
            t += arr[j]==1
            if t == k:
                break
            j -= 1

        while i < n:
            if arr[i]: break
            i += 1

        for l in range(j, n):
            if arr[i] != arr[l]:
                return [-1, -1]
            i += 1
        
        for k in range(i, n):
            if arr[k]: break
        
        for l in range(j, n):
            if arr[k] != arr[l]:
                return [-1, -1]
            k += 1
            
        return [i - 1, k]