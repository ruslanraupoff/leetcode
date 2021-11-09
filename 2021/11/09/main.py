class Solution:
    def findNumOfValidWords(self, words: List[str], puzzles: List[str]) -> List[int]:
        mp = {}
        def mask(w):
            res = 0
            for l in w:
                res |= 1 << (ord(l)-ord('a'))
            return res
        
        for w in words:
            k = mask(w)
            if k in mp:
                mp[k] += 1
            else:
                mp[k] = 1

        res = [0] * len(puzzles)
        for i, p in enumerate(puzzles):
            fb = 1 << (ord(p[0]) - ord('a'))
            ob = mask(p)
            sb = ob
            while sb != 0:
                if sb & fb != 0:
                    if sb in mp:
                        res[i] += mp[sb]
                sb = (sb - 1) & ob

        return res
