class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        next = [[]] * numCourses
        pre = [0] * numCourses

        for r in prerequisites:
            next[r[1]].append(r[0])
            pre[r[0]] += 1
        
        i, res = 0, [0] * numCourses
        while i < numCourses:
            j = 0
            while j < numCourses:
                if pre[j] == 0:
                    break
                j += 1
            if j == numCourses:
                return None
            pre[j] = -1
            for c in next[j]:
                pre[c] -= 1
            res[i] = j
            i += 1

        return res