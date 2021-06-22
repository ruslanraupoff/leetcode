class Solution:
    def numMatchingSubseq(self, s: str, words: List[str]) -> int:
        r, hs, uhs = 0, set(), set()
        def is_sub(word):
            idx = -1
            for w in word:
                idx = s.find(w, idx + 1)
                if idx == -1:
                    return False
            return True
        for w in words:
            if hs.__contains__(w):
                r += 1
                continue
            if uhs.__contains__(w): continue
            if is_sub(w):
                hs.add(w)
                r += 1
            else:
                uhs.add(w)
        return r