class Solution:
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        hm = {}
        res = []
        for i in range(len(words)):
            hm[words[i]] = i

        def ispalindrome(x):
            return x == x[::-1]
        
        for i in range(len(words)):
            word = words[i]
            n = len(word)
            for j in range(n+1):
                pre = word[:j]
                suf = word[j:]
                if ispalindrome(pre):
                    rstr = suf[::-1]
                    if rstr != word and rstr in hm:
                        res.append([hm[rstr], i])
                if j != n and ispalindrome(suf):
                    rstr = pre[::-1]
                    if rstr != word and rstr in hm:
                        res.append([i, hm[rstr]])
        return res