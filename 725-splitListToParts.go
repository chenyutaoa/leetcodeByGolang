package leetcode

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := getLen(head)
	result := make([]*ListNode,length/k+1)

	return result
}

func getLen(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}

	return count
}
