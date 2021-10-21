class RandomizedSet:

    def __init__(self):
        self.a = []
        self.m = {}

    def insert(self, val: int) -> bool:
        if val in self.m:
            return False
        
        self.a.append(val)
        self.m[val] = len(self.a) - 1
        return True

    def remove(self, val: int) -> bool:
        if val not in self.m:
            return False

        self.a[self.m[val]] = self.a[len(self.a)-1]
        self.m[self.a[len(self.a)-1]] = self.m[val]
        self.a = self.a[:len(self.a)-1]
        del self.m[val]
        return True
        

    def getRandom(self) -> int:
        return self.a[random.randrange(len(self.a))]
        


# Your RandomizedSet object will be instantiated and called as such:
# obj = RandomizedSet()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()