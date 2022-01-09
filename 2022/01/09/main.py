class Solution:
    def isRobotBounded(self, instructions: str) -> bool:
        step = [(0,1), (-1,0), (0,-1), (1,0)]
        c, x, y = 0, 0, 0
        for i in instructions:
            if i=="G":
                x += step[c][0]
                y += step[c][1]
            elif i == 'L':
                c = (c - 1) % 4
            elif i == 'R':
                c = (c + 1) % 4
        return (x==0 and y==0) or c != 0