package leetcode

//给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//  2 4 3
//+ 5 6 4
//= 7 0 8

//  2 4 6
//+ 5 6 4
//= 7 0 0 1
type List struct {
	headNode *ListNode //头节点
	curNode  *ListNode //当前节点
}

func (this *List) ISEmpty() bool {

	//判断单链表是否为空，只需要判断头节点是否为空即可
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	s := 0
//	l3 := &List{}
//	for l1 != nil || l2 != nil {
//		v1, v2, v3 := 0, 0, 0
//		if l1 != nil {
//			v1 = l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			v2 = l2.Val
//			l2 = l2.Next
//		}
//		if v1+v2+s >= 10 {
//			v3 = (v1 + v2 + s) % 10
//			s = 1
//		} else {
//			v3 = v1 + v2 + s
//			s = 0
//		}
//		if l3.ISEmpty() {
//			l3 = &List{headNode: &ListNode{
//				Val:  v3,
//				Next: nil,
//			}}
//			l3.curNode = l3.headNode
//		} else {
//			l3.curNode.Next = &ListNode{
//				Val:  v3,
//				Next: nil,
//			}
//			l3.curNode = l3.curNode.Next
//		}
//	}
//	if s == 1 {
//		l3.curNode.Next = &ListNode{
//			Val:  1,
//			Next: nil,
//		}
//	}
//	return l3.headNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s := 0
	l3 := new(ListNode)
	head:= l3
	for l1 != nil || l2 != nil {
		v1, v2, v3 := 0, 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		if v1+v2+s >= 10 {
			v3 = (v1 + v2 + s) % 10
			s = 1
		} else {
			v3 = v1 + v2 + s
			s = 0
		}
		l3.Next = new(ListNode)
		l3 = l3.Next
		l3.Val = v3
	}
	if s == 1 {
		l3.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return head.Next
}
