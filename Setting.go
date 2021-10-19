package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
type TrieNode struct {
	isWord       bool
	nextTrieNode [26]*TrieNode
}
