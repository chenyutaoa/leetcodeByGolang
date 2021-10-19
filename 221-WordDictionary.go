package leetcode

type WordDictionary struct {
	wordNde *TrieNode
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{wordNde: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	next := this.wordNde
	for _, s := range word {
		if next.nextTrieNode[s-'a'] == nil{
			next.nextTrieNode[s-'a'] = &TrieNode{}
		}
		next = next.nextTrieNode[s-'a']
	}
	next.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(index int, node *TrieNode) bool {
		if index == len(word) {
			return node.isWord
		}
		ch := word[index]
		if ch != '.' {
			child := node.nextTrieNode[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.nextTrieNode {
				child := node.nextTrieNode[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, this.wordNde)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
