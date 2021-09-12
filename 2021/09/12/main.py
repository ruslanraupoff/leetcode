class Solution:
    def reachableNodes(self, edges: List[List[int]], maxMoves: int, n: int) -> int:
        flag = True
        steps = [-1] * n
        steps[0] = maxMoves
        while(flag):
            flag = False
            for edge in edges:
                if steps[edge[0]] > 0:
                    na = steps[edge[0]] - edge[2] - 1
                    if (na > steps[edge[1]]):
                        steps[edge[1]] = na
                        flag = True
                if steps[edge[1]] > 0:
                    na = steps[edge[1]] - edge[2] - 1
                    if (na > steps[edge[0]]):
                        steps[edge[0]] = na
                        flag = True
        
        res = 0
        for i in range(n):
            if steps[i] != -1:
                res+=1
        for edge in edges:
            tmp = 0
            if steps[edge[0]] > 0:
                tmp += steps[edge[0]]
            if steps[edge[1]] > 0:
                tmp += steps[edge[1]]
            
            if tmp > edge[2]:
                res += edge[2]
            else:
                res += tmp
        
        return res