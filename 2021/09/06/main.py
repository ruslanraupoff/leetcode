class Solution:
    def slowestKey(self, releaseTimes: List[int], keysPressed: str) -> str:
        i, r, mx = 0, '', -1
        while i < len(keysPressed):
            k = keysPressed[i]
            if i == 0:
                mx = releaseTimes[i]
                r = k
            else:
                d = releaseTimes[i] - releaseTimes[i - 1]
                if d > mx:
                    mx = d
                    r = k
                elif d == mx and k > r:
                    r = k
            i += 1
        return r