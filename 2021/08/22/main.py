class Solution:
    def rectangleArea(self, rectangles: List[List[int]]) -> int:
        rect = []
        for x0, y0, x1, y1 in rectangles:
            if x0 == x1: continue
            rect.append((x0, 1, y0, y1))
            rect.append((x1, 0, y0, y1))
        rect.sort()
        l = rect[0][0]
        r = 0
        e = []
        for t, state, y0, y1 in rect:
            r += self.calc_y_len(e) * (t - l)
            if state:
                e.append((y0, y1))
                e.sort()
            else:
                e.remove((y0, y1))
            l = t
        return r % (10 ** 9 + 7)

    def calc_y_len(self, e):
        r, c = 0, -1
        for y0, y1 in e:
            c = max(c, y0)
            r += max(0, y1 - c)
            c = max(c, y1)
        return r