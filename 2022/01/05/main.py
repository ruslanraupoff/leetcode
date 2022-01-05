class Solution:
    def partition(self, s: str) -> List[List[str]]:
        result = []
        current = []*len(s)
        def pal(s):
            if len(s) <= 1:
                return True
            a, b = 0, len(s)-1
            while a < b:
                if s[a] != s[b]:
                    return False
                a+=1
                b-=1
            return True
        def dfs(i, cur, result):
            if i == len(s):
                tmp = cur.copy()
                result.append(tmp)
                return
            for j in range(i, len(s)):
                if pal(s[i:j+1]):
                    t = cur.copy()
                    t.append(s[i:j+1])
                    dfs(j+1, t, result)

        dfs(0, current, result)
        return result
