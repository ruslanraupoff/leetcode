class Solution:
    def findDuplicate(self, paths: List[str]) -> List[List[str]]:
        d = {}
        for p in paths:
            ps = p.split(" ")
            for f in ps[1:]:
                fs = f.split("(")
                n, c = fs[0], fs[1][:-1]
                if not c in d:
                    d[c] = []
                d[c].append(ps[0] + "/" + n)
        return [p for p in d.values() if len(p) > 1]