class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        boxTypes.sort(key=lambda x: x[1], reverse=True)
        cnt = 0
        res = 0
        for box in boxTypes:
            for num in range(box[0]):
                if cnt < truckSize:
                    res += box[1]
                    cnt += 1
        return res