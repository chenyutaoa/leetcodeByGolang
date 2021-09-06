package leetcode

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{
			head: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
								Next: &ListNode{
									Val:  8,
									Next: &ListNode{
										Val:  0,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			k: 1,
		}, &ListNode{
			Val:  0,
			Next: nil,
		}},
		{"", args{
			head: &ListNode{
				Val: 1,
				Next: nil,
			},
			k: 1,
		}, &ListNode{
			Val: 1,
			Next: nil,
		}},
		{"", args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: nil,
				},
			},
			k: 2,
		}, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: nil,
			},
		}},
		{"", args{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: nil,
				},
			},
			k: 1,
		}, &ListNode{
				Val: 2,
				Next: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
