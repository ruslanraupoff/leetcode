class Solution:
    def solve(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        m, n = len(board), len(board[0])
        if m <= 2 or n <= 2:
            return

        isEscaped = [[0] * n for _ in range(m)]
        self.idxM  = []
        self.idxN  = []

        def bfs(i, j):
            isEscaped[i][j] = True
            self.idxM .append(i)
            self.idxN .append(j)

            while len(self.idxM ) > 0:
                i = self.idxM [0]
                j = self.idxN [0]
                self.idxM  = self.idxM [1:]
                self.idxN  = self.idxN [1:]

                if 0 <= i-1 and board[i-1][j] == 'O' and not isEscaped[i-1][j]:
                    self.idxM .append(i-1)
                    self.idxN .append(j)
                    isEscaped[i-1][j] = True
                

                if 0 <= j-1 and board[i][j-1] == 'O' and not isEscaped[i][j-1]:
                    self.idxM .append(i)
                    self.idxN .append(j-1)
                    isEscaped[i][j-1] = True
                

                if i+1 < m and board[i+1][j] == 'O' and not isEscaped[i+1][j]:
                    self.idxM .append(i+1)
                    self.idxN .append(j)
                    isEscaped[i+1][j] = True
                

                if j+1 < n and board[i][j+1] == 'O' and not isEscaped[i][j+1]:
                    self.idxM .append(i)
                    self.idxN .append(j+1)
                    isEscaped[i][j+1] = True

        for j in range(n):
            if board[0][j] == 'O' and not isEscaped[0][j]:
                bfs(0, j)
            
            if board[m-1][j] == 'O' and not isEscaped[m-1][j]:
                bfs(m-1, j)
            
        
        for i in range(m):
            if board[i][0] == 'O' and not isEscaped[i][0]:
                bfs(i, 0)
            
            if board[i][n-1] == 'O' and not isEscaped[i][n-1]:
                bfs(i, n-1)
            
        

        for i in range(1, m-1):
            for j in range(1, n-1):
                if board[i][j] == 'O' and not isEscaped[i][j]:
                    board[i][j] = 'X'
                
            
        