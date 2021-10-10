
class TrieNode():
    def __init__(self):
        self.next = {}
        self.word = ""

class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        r, n, m = [], len(board[0]), len(board)
        
        trie = TrieNode()
        for word in words:
            cur = trie
            for c in word:
                if c not in cur.next:
                    cur.next[c] = TrieNode()
                cur = cur.next[c]
            cur.word = word

        def dfs(i, j, trie):
            c = board[i][j]
            if c == '#' or c not in trie.next:
                return

            trie = trie.next[c]
            if len(trie.word) > 0:
                r.append(trie.word)
                trie.word = ""

            board[i][j] = '#'
            if i > 0:
                dfs(i-1, j, trie)

            if i < len(board)-1:
                dfs(i+1, j, trie)

            if j > 0:
                dfs(i, j-1, trie)

            if j < len(board[0])-1:
                dfs(i, j+1, trie)

            board[i][j] = c

        for i in range(m):
            for j in range(n):
                dfs(i, j, trie)

        return r
