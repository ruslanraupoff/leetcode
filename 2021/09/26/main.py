class Solution:
    def movesToChessboard(self, board: List[List[int]]) -> int:
        n = len(board)
        for i in range(n):
            for j in range(n):
                if board[0][0]^board[i][0]^board[0][j]^board[i][j] == 1:
                    return -1
        
        rsm, csm, rsp, csp = 0, 0, 0, 0

        for i in range(n):
            rsm += board[0][i]
            csm += board[i][0]
            if board[i][0] == i%2:
                rsp += 1
            if board[0][i] == i%2:
                csp += 1

        if rsm < n/2 or (n+1)/2 < rsm:
            return -1

        if csm < n/2 or (n+1)/2 < csm:
            return -1

        if n%2 == 1:
            if csp%2 == 1:
                csp = n - csp
            if rsp%2 == 1:
                rsp = n - rsp
        else:
            csp = min(n-csp, csp)
            rsp = min(n-rsp, rsp)

        return (csp + rsp) // 2