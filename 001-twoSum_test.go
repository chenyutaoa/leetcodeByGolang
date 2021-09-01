package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{nums: []int{0, 1, 2, 3, 4, 5}, target: 1}, []int{0, 1}},
		{"", args{nums: []int{5, 6, 7, 8, 9}, target: 12}, []int{0, 2}},
		{"", args{nums: []int{2, 7, 11, 15}, target: 9}, []int{0, 1}},
		{"", args{nums: []int{3, 2, 4}, target: 6}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
			if got := twoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
