package main

type TrieNode struct {
	next [26]*TrieNode
	word string
}

func findWords(board [][]byte, words []string) []string {
	var r []string

	m := len(board)
	if m == 0 {
		return r
	}

	n := len(board[0])
	if n == 0 {
		return r
	}

	trie := buildTrie(words)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, trie, &r)
		}
	}

	return r
}

func dfs(board [][]byte, i int, j int, trie *TrieNode, r *[]string) {
	c := board[i][j]
	if c == '#' || trie.next[int(c-'a')] == nil {
		return
	}

	trie = trie.next[int(c-'a')]
	if len(trie.word) > 0 {
		// Found one
		*r = append(*r, trie.word)
		trie.word = ""
	}

	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, trie, r)
	}

	if i < len(board)-1 {
		dfs(board, i+1, j, trie, r)
	}

	if j > 0 {
		dfs(board, i, j-1, trie, r)
	}

	if j < len(board[0])-1 {
		dfs(board, i, j+1, trie, r)
	}

	board[i][j] = c
}

func buildTrie(words []string) *TrieNode {
	root := new(TrieNode)
	for _, word := range words {
		cur := root
		for _, c := range word {
			cidx := int(c - 'a')
			if cur.next[cidx] == nil {
				cur.next[cidx] = new(TrieNode)
			}
			cur = cur.next[cidx]
		}
		cur.word = word
	}

	return root
}
