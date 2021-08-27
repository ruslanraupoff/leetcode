class Solution:
    def findLUSlength(self, strs: List[str]) -> int:
        def is_sub(s: str, t: str) -> bool:
            i = 0
            for ch in t:
                if i < len(s) and s[i] == ch:
                    i += 1
            return i == len(s)
        strs.sort(key=len, reverse=True)
        for i, word1 in enumerate(strs):
            if all(not is_sub(word1, word2) for j, word2 in enumerate(strs) if i != j):
                return len(word1)
        return -1