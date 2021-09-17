package leetcode

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{
			board: [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			words: []string{"oath", "pea", "eat", "rain"},
		}, []string{"oath", "eat"}},
		{"", args{
			board: [][]byte{{'a', 'b'}, {'c', 'd'}},
			words: []string{"abcb", "abd", "eat", "rain"},
		}, []string{"abd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
