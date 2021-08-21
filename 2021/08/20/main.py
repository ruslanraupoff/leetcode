class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        return self.isValidRow(board) and self.isValidCol(board) and self.isValidSquare(board)

    def isValidRow(self, board):
        for row in board:
            if not self.isValid(row): return False
        return True

    def isValidCol(self, board):
        for col in zip(*board):
            if not self.isValid(col): return False
        return True

    def isValidSquare(self, board):
        for i in (0, 3, 6):
            for j in (0, 3, 6):
                s = [board[x][y] for x in range(i,i+3) for y in range(j,j+3)]
                if not self.isValid(s): return False
        return True

    def isValid(self, nums):
        res = [num for num in nums if num != '.']
        return len(res) == len(set(res))