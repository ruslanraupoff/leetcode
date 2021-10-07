class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        k, m, n = 0, len(board), len(board[0])

        def dfs(r, c, k):
            if len(word) == k:
                return True
            
            if r < 0 or m <= r or c < 0 or n <= c or board[r][c] != word[k]:
                return False

            t = board[r][c]
            board[r][c] = 0

            if dfs(r-1, c, k+1) or dfs(r+1, c, k+1) or dfs(r, c-1, k+1) or dfs(r, c+1, k+1):
                return True

            board[r][c] = t
            return False

        for i in range(m):
            for j in range(n):
                if dfs(i, j, 0):
                    return True

        return False