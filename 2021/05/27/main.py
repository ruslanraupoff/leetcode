class Solution:
    def maxProduct(self, words: List[str]) -> int:
        r = 0
        for i in range(len(words) - 1):
            a = words[i]
            la = len(a)
            for j in range(i+1, len(words)):
                b = words[j]
                lb = len(b)
                for w in a:
                    if w in b:
                        lb = 0
                        break
                r = max(r, la * lb)
        return r