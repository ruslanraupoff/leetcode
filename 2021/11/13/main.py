class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        res = [0] * n
        stack = [0] * n
        top = -1
        for i in range(n):
            while top >= 0 and temperatures[stack[top]] < temperatures[i]:
                res[stack[top]] = i - stack[top]
                top -= 1

            top += 1
            stack[top] = i

        return res