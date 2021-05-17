class Solution:
    def longestStrChain(self, words: List[str]) -> int:
        self.mp = {}
        self.ans = 0
        for w in words:
            self.mp[w] = 1
        words.sort(key=len)
        for w in words:
            for i in range(len(w)):
                p = str(w[:i] + w[i+1:])
                if p in self.mp and self.mp[p] + 1 > self.mp[w]:
                    self.mp[w] = self.mp[p] + 1
            if self.mp[w] > self.ans:
                self.ans = self.mp[w]
        return self.ans
        