class Solution:
    def reverseOnlyLetters(self, s: str) -> str:
        l, r = 0, len(s) - 1
        b = list(s)
        def isLetter(c):
            return ('a' <= c and c <= 'z') or ('A' <= c and c <= 'Z')
        while l < r:
            while l < r and not isLetter(b[l]):
                l += 1
            
            while l < r and not isLetter(b[r]):
                r -= 1
            
            b[l], b[r] = b[r], b[l]
            l += 1
            r -= 1
    
        return ''.join(b)

