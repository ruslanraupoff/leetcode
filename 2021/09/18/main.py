class Solution:
    def addOperators(self, num: str, target: int) -> List[str]:
        res = []
        def integer(s):
            r = ord(s[0]) - ord('0')
            for i in range(1, len(s)):
                r = 10 * r + ord(s[i]) - ord('0')
            return r

        def dfs(s, t, r, p):
            if len(s) == 0:
                if r == target:
                    res.append(t)
                return
            
            for i in range(1, len(s)+1):
                cs = s[:i]
                if cs[0] == '0' and len(cs) > 1:
                    return
                cn = integer(cs)
                ns = s[i:]

                if len(t) == 0:
                    dfs(ns, cs, cn, cn)
                else:
                    dfs(ns, t+"+"+cs, r+cn, cn)
                    dfs(ns, t+"-"+cs, r-cn, -cn)
                    dfs(ns, t+"*"+cs, r-p+p*cn, p*cn)

        dfs(num, "", 0, 0)
        
        return res
