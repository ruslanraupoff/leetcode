import heapq

class MedianFinder:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.left, self.right = [], []

    def addNum(self, num: int) -> None:
        if len(self.left) == len(self.right):
            heapq.heappush(self.left, -num)
            tmp = heapq.heappop(self.left)
            heapq.heappush(self.right, -tmp)
        else:
            heapq.heappush(self.right, num)
            tmp = heapq.heappop(self.right)
            heapq.heappush(self.left, -tmp)

    def findMedian(self) -> float:
        if len(self.left) == len(self.right):
            return (self.right[0]-self.left[0]) / 2
        
        return self.right[0]


# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
