import heapq

class Solution:
    def scheduleCourse(self, courses: List[List[int]]) -> int:
        hp = []
        courses.sort(key=lambda x: x[1])
        cur = 0
        for a, b in courses:
            cur += a
            heapq.heappush(hp, -a)
            if cur > b:
                cur += heapq.heappop(hp)
        return len(hp)

if __name__ == '__main__':
    print('__main__')
    courses = [[3,2],[4,3]]#[[100,200],[200,1300],[1000,1250],[2000,3200]]
    res = Solution().scheduleCourse(courses)
    print(res)
