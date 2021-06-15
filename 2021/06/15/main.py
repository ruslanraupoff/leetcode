class Solution:
    def makesquare(self, matchsticks: List[int]) -> bool:
        n = len(matchsticks)
        per = sum(matchsticks)
        side = per // 4
        if side * 4 != per:
            return False

        matchsticks.sort(reverse=True)

        sq = [0]*4

        def dfs(idx):
            if idx == n:
                return sq[0] == sq[1] == sq[2] == side

            for i in range(4):
                if sq[i] + matchsticks[idx] <= side:
                    sq[i] += matchsticks[idx]

                    if dfs(idx+1):
                        return True
                    sq[i] -= matchsticks[idx]

            return False

        if not matchsticks:
            return False

        return dfs(0)
