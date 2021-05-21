class Solution:
    def findAndReplacePattern(self, words: List[str], pattern: str) -> List[str]:
        def ispattern(w):
            return len(set(w)) == len(set(pattern)) == len(set(zip(w, pattern)))
        return [w for w in words if ispattern(w)]