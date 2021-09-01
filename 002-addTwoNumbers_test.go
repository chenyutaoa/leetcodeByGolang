package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{l1: &ListNode{Val: 1, Next: &ListNode{
			Val: 1,
			Next: &ListNode{Val: 1, Next: &ListNode{
				Val: 1,
				Next: &ListNode{Val: 1, Next: &ListNode{
					Val:  1,
					Next: nil,
				}},
			}},
		}}, l2: &ListNode{Val: 1, Next: &ListNode{
			Val: 1,
			Next: &ListNode{Val: 1, Next: &ListNode{
				Val: 1,
				Next: &ListNode{Val: 1, Next: &ListNode{
					Val:  1,
					Next: nil,
				}},
			}},
		}}}, &ListNode{Val: 2, Next: &ListNode{
			Val: 2,
			Next: &ListNode{Val: 2, Next: &ListNode{
				Val: 2,
				Next: &ListNode{Val: 2, Next: &ListNode{
					Val:  2,
					Next: nil,
				}},
			}},
		}}},
		{"", args{l1: &ListNode{Val: 1, Next: &ListNode{
			Val:  1,
			Next: &ListNode{Val: 1},
		}}, l2: &ListNode{Val: 1}},
			&ListNode{Val: 2, Next: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 1},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
