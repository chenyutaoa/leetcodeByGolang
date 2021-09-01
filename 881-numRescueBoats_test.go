package leetcode

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{
			people: []int{1, 2},
			limit:  3,
		}, 1},
		{"", args{
			people: []int{3, 2, 2, 1},
			limit:  3,
		}, 3},
		{"", args{
			people: []int{3, 5, 3, 4},
			limit:  5,
		}, 4},
		{"", args{
			people: []int{21, 40, 16, 24, 30},
			limit:  50,
		}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
