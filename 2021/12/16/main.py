class Solution:
    def findMinHeightTrees(self, n: int, edges: List[List[int]]) -> List[int]:
        if n == 1:
            return [0]

        ns = [[] for _ in range(n)]
        cn = [0] * n
        for e in edges:
            ns[e[0]].append(e[1])
            ns[e[1]].append(e[0])
            cn[e[0]]+=1
            cn[e[1]]+=1

        ls = []
        for i in range(n):
            if cn[i] == 1:
                ls.append(i)

        while n > 2:
            nls = []
            for leaf in ls:
                n-=1
                for nd in ns[leaf]:
                    cn[nd]-=1
                    if cn[nd] == 1:
                        nls.append(nd)
            ls = nls

        res = []
        for v in ls:
            res.append(v)

        return res