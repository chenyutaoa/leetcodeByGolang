package leetcode

import (
	"reflect"
	"testing"
)

func Test_rand10(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
		{"",args{},1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rand10(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rand10() = %v, want %v", got, tt.want)
			}
		})
	}
}
