class Solution:
    def canReach(self, arr: List[int], start: int) -> bool:
        
        if start < 0 or start > len(arr)-1 or arr[start] == -1:
            return False
        
        delta = arr[start]
        arr[start] = -1

        return  (delta == 0) or self.canReach(arr, start+delta) or self.canReach(arr, start-delta)