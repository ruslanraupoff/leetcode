class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        self.solve(board, 0)
    
    def solve(self, board, k):
        if k == 81:
            return True

        r, c = k//9, k%9
        
        if board[r][c] != '.':
            return self.solve(board, k+1)

        bi, bj = r//3*3, c//3*3

        def isValid(b):
            for n in range(9):
                if board[r][n] == b or board[n][c] == b or board[bi+n//3][bj+n%3] == b:
                    return False
            return True
        

        for b in "123456789":
            if isValid(b):
                board[r][c] = b
                if self.solve(board, k+1):
                    return True

        board[r][c] = '.'

        return False