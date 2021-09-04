class Solution:
    def sumOfDistancesInTree(self, n: int, edges: List[List[int]]) -> List[int]:
        tree = collections.defaultdict(set)
        res = [0] * n
        count = [0] * n
        for i, j in edges:
            tree[i].add(j)
            tree[j].add(i)

        def dfs(root=0, seen=set()):
            seen.add(root)
            for i in tree[root]:
                if i not in seen:
                    dfs(i, seen)
                    count[root] += count[i]
                    res[root] += res[i] + count[i]
            count[root] += 1

        def dfm(root=0, seen=set()):
            seen.add(root)
            for i in tree[root]:
                if i not in seen:
                    res[i] = res[root]+ n - 2 * count[i]
                    dfm(i, seen)
        dfs()
        dfm()
        return res