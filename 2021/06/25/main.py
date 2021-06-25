class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        i, n = 0, len(edges)
        p = [i for i in range(n+1)]
        def ifind(i):
            while i != p[i]:
                i = p[i]
            return i
        for f, t in edges:
            ff = ifind(f)
            ft = ifind(t)
            if ff == ft:
                break
            p[ff] = ft
            i += 1
        return edges[i]

