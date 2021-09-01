package leetcode

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "hello"}, "holle"},
		{"", args{s: "leetcode"}, "leotcede"},
		{"", args{s: "aeiou"}, "uoiea"},
		{"", args{s: "Euston saw I was not Sue."}, "euston saw I was not SuE."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseVowels11(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "hello"}, "holle"},
		{"", args{s: "leetcode"}, "leotcede"},
		{"", args{s: "aeiou"}, "uoiea"},
		{"", args{s: "Euston saw I was not Sue."}, "euston saw I was not SuE."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels1(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels1() = %v, want %v", got, tt.want)
			}
		})
	}
}
