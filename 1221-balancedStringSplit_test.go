package leetcode

import "testing"

func Test_balancedStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"",args{s: ""},0},
		{"",args{s: "RLRRLLRLRL"},4},
		{"",args{s: "RRRLLL"},1},
		{"",args{s: "RLLLLRRRLR"},3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
