package leetcode

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{arr: []int{0, 1, 0}}, 1},
		{"", args{arr: []int{1, 3, 5, 4, 2}}, 2},
		{"", args{arr: []int{0, 10, 5, 2}}, 1},
		{"", args{arr: []int{3, 4, 5, 1}}, 2},
		{"", args{arr: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}}, 2},
		{"", args{arr: []int{12,13,19,41,55,69,70,71,96,72}}, 8},
		{"", args{arr: []int{55,59,63,99,97,94,84,81,79,66,40,38,33,23,22,21,17,9,7}}, 3},
		{"", args{arr: []int{8,18,24,31,37,42,43,56,65,73,93,98,100,98,76,72,69,24,23}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
