class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        i, m = 0, len(flowerbed)
        while n > 0 and i < m:
            if flowerbed[i] == 1:
                i += 2
            elif i < m - 1 and flowerbed[i + 1] == 1:
                i += 3
            else:
                n -= 1
                i += 2

        return n == 0