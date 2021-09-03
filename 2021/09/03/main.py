class Solution:
    def outerTrees(self, trees: List[List[int]]) -> List[List[int]]:
       
        def is_clockwise(p0: List[int], p1: List[int], p2: List[int]) -> bool:
            return (p1[1] - p0[1]) * (p2[0] - p0[0]) > (p2[1] - p0[1]) * (p1[0] - p0[0])

        strees = sorted(trees)

        r = []
        for p in strees:
            while len(r) > 1 and is_clockwise(r[-2], r[-1], p):
                r.pop()

            r.append(p)

        if len(r) == len(trees):
            return r

        r.pop()
        for p in reversed(strees):
            while len(r) > 1 and is_clockwise(r[-2], r[-1], p):
                r.pop()

            r.append(p)

        r.pop()

        return r
