class Solution:
    def tictactoe(self, moves: List[List[int]]) -> str:
        a, b = 1, -1
        player = a
        rows, cols = [0] * 3, [0] * 3
        up, down = 0, 0

        for r, c in moves:
            rows[r] += player
            cols[c] += player
            if r == c:
                up += player
            if r + c == 2:
                down += player

            player = -player

        lines = rows + cols + [up, down]
        if any(line == 3 * a for line in lines):
            return "A"
        if any(line == 3 * b for line in lines):
            return "B"
        return "Draw" if len(moves) == 9 else "Pending"

