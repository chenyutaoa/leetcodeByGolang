package leetcode

import "testing"

func Test_chalkReplacer(t *testing.T) {
	type args struct {
		chalk []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"",args{
			chalk: []int{5, 1, 5},
			k:     22,
		},0},
		{"",args{
			chalk: []int{3,4,1,2},
			k:     25,
		},1},
		{"",args{
			chalk: []int{3,4,1,2},
			k:     7,
		},2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chalkReplacer(tt.args.chalk, tt.args.k); got != tt.want {
				t.Errorf("chalkReplacer() = %v, want %v", got, tt.want)
			}
		})
	}
}
