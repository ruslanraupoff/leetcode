class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort()
        res = []
        temp = intervals[0]
        for i in range(len(intervals)):
            if intervals[i][0] <= temp[1]:
                temp[1] = max(temp[1], intervals[i][1])
            else:
                res.append(temp)
                temp = intervals[i]
        
        res.append(temp)
        return res
