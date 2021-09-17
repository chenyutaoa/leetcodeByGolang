package leetcode

type trie struct {
	children [26]*trie
	word     string
}

func (t *trie) insert(word string) {
	for _, c := range word {
		c -= 'a'
		if t.children[c] == nil {
			t.children[c] = &trie{}
		}
		t = t.children[c]
	}
	t.word = word
}

func findWords(board [][]byte, words []string) []string {
	t := &trie{}
	for _, word := range words {
		t.insert(word)
	}
	m, n := len(board), len(board[0])
	seen := map[string]bool{}
	var dfs func(tr *trie, x, y int)
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(t *trie, x, y int) {
		c := board[x][y]
		node := t.children[c-'a']
		if node == nil {
			return
		}
		if node.word != "" {
			seen[node.word] = true
		}
		board[x][y] = '!'
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '!' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = c
	}

	for x, bo := range board {
		for y, _ := range bo {
			dfs(t, x, y)
		}
	}
	var res []string
	for word, _ := range seen {
		res = append(res, word)
	}
	return res
}
