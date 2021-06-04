class Solution:
    def openLock(self, deadends: List[str], target: str) -> int:
        ss = set(deadends)
        
        if '0000' in ss:
            return -1
        
        if target == '0000':
            return 0
        
        queue = ['0000']
        r, ms = 0, set(queue)
        
        def conv(s):
            for i in range(len(s)):
                c = int(s[i])
                yield s[:i] + str((c+1)%10) + s[i+1:]
                yield s[:i] + str((c-1)%10) + s[i+1:]

        while queue:
            r += 1
            qn = []
            for i in queue:
                for j in conv(i):
                    if j == target:
                        return r
                    if j in ss:
                        continue
                    if j not in ms:
                        qn.append(j)
                    ms.add(j)
            queue = qn
        return -1