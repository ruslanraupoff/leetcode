class Solution:
    def customSortString(self, order: str, str: str) -> str:
        res = ''
        for i in range(len(order)):
            o = order[i:i+1]
            count = str.count(o)
            res += o * count
            str = str.replace(o, "", -1)
        return res + str