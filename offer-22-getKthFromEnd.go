package leetcode

//func getKthFromEnd(head *ListNode, k int) *ListNode {
//	if head.Next == nil {
//		return head
//	}
//	var nHead *ListNode
//	for head != nil {
//		h := &ListNode{Val: head.Val}
//		h.Next = nHead
//		nHead = h
//		head = head.Next
//	}
//	i := 0
//	for i != k {
//		if nHead == nil {
//			return nil
//		}
//		h := &ListNode{Val: nHead.Val}
//		h.Next = head
//		head = h
//		nHead = nHead.Next
//		i++
//	}
//	return head
//}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	s, e := head, head
	for i := 0; i < k; i++ {
		e = e.Next
	}
	for e != nil {
		s = s.Next
		e = e.Next
	}
	return s
}
