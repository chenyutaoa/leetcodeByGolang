package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{
			s:       "PAYPALISHIRING",
			numRows: 3,
		}, "PAHNAPLSIIGYIR"},
		{"", args{
			s:       "PAYPALISHIRING",
			numRows: 4,
		}, "PINALSIGYAHRPI"},
		{"", args{
			s:       "AB",
			numRows: 1,
		}, "AB"},
		{"", args{
			s:       "ABc",
			numRows: 2,
		}, "AB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
