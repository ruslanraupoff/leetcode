class Solution:
    def frequencySort(self, s: str) -> str:
        l, mp = [], {}
        for c in s:
            mp[c] = mp.get(c, 0) + 1
        for k in mp:
            if mp[k] == 0:
                continue
            l.append(k * mp[k])
        
        l = reversed(sorted(l, key=len))
        
        return ''.join(l)