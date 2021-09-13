package leetcode

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{s: "()"}, true},
		{"", args{s: ")("}, false},
		{"", args{s: "((*()"}, false},
		{"", args{s: "()))"}, false},
		{"", args{s: "(*)"}, true},
		{"", args{s: "(*))"}, true},
		{"", args{s: "((())"}, false},
		{"", args{s: "(*))()"}, true},
		{"", args{s: "**("}, false},
		{"", args{s: "(**"}, true},
		{"", args{s: "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"}, false},
		{"", args{s: "(((()*())))((()(((()(()))()**(*)())))())()()*"}, true},
		{"", args{s: "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
