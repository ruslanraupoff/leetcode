class CombinationIterator:

    def __init__(self, characters: str, combinationLength: int):
        self.index = 0
        self.combines = []
        
        def dfs(characters, start, combinationLength, currStr, res):
            if len(currStr) == combinationLength:
                res.append(currStr)
                return
            for i in range(start, len(characters)):
                currStr += characters[i:i+1]
                dfs(characters, i+1, combinationLength, currStr, res)
                currStr = currStr[:len(currStr)-1]
        
        dfs(characters, 0, combinationLength, "", self.combines)
    
        

    def next(self) -> str:
        val = self.combines[self.index]
        self.index += 1
        return val
        

    def hasNext(self) -> bool:
        return len(self.combines) != self.index
        
# Your CombinationIterator object will be instantiated and called as such:
# obj = CombinationIterator(characters, combinationLength)
# param_1 = obj.next()
# param_2 = obj.hasNext()

