class MapSum:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.mp = {}
        

    def insert(self, key: str, val: int) -> None:
        self.mp[key] = val

    def sum(self, prefix: str) -> int:
        s, l = 0, len(prefix)
        for k, v  in self.mp.items():
            if k[:l] == prefix:
                s += v
        return s


# Your MapSum object will be instantiated and called as such:
# obj = MapSum()
# obj.insert(key,val)
# param_2 = obj.sum(prefix)