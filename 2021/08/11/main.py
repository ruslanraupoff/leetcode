class Solution:
    def canReorderDoubled(self, arr: List[int]) -> bool:
        mp = {}
        arr.sort(key=lambda x: abs(x))
        for n in arr:
            mp[n] = mp.get(n, 0) + 1
        for n in arr:
            if mp[n] <= 0:
                continue
            if 2 * n in mp and mp[2 * n] > 0:
                mp[n] -= 1
                mp[2 * n] -= 1
            else:
                return False
        return True