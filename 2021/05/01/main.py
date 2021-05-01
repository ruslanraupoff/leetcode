class WordFilter:

    def __init__(self, words: List[str]):
        self.mp = {}
        for i in range(len(words)):
            w = words[i]
            for j in range(1, len(w) + 1):
                for k in range(1, len(w) + 1):
                    st = w[0:j]+"#"+w[k-len(w):]
                    self.mp[st] = i

    def f(self, prefix: str, suffix: str) -> int:
        st = prefix + "#" + suffix
        return self.mp[st] if st in self.mp else -1
        

# Your WordFilter object will be instantiated and called as such:
# obj = WordFilter(words)
# param_1 = obj.f(prefix,suffix)