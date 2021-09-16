package leetcode

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	for _, c := range word {
		c -= 'a'
		if t.children[c] == nil {
			t.children[c] = &Trie{}
		}
		t = t.children[c]
	}
	t.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	for _, c := range word {
		c -= 'a'
		if t.children[c] == nil {
			return false
		}
		t = t.children[c]
	}
	return t.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		c -= 'a'
		if t.children[c] == nil {
			return false
		}
		t = t.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := ConstructorTrie();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
