class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        c = range(n)
        l = []
        for per in permutations(c):
            if n == len(set(per[i]+i for i in c)) == len(set(per[i]-i for i in c)):
                st = []
                for i in range(len(per)):
                    s, p = '', per[i]
                    for i in c:
                        s += 'Q' if i == p else '.'
                    st.append(s)
                l.append(st)
        return l
    