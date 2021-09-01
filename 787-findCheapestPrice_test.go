package leetcode

import "testing"

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       1,
		}, 200},
		{"", args{
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       0,
		}, 500},
		{"", args{
			n:       5,
			flights: [][]int{{0, 1, 100}, {1, 4, 100}, {2, 4, 500}, {1, 3, 20}, {3, 4, 10}},
			src:     0,
			dst:     4,
			k:       3,
		}, 130},
		{"", args{
			n:       5,
			flights: [][]int{{0, 1, 100}, {1, 4, 100}, {2, 4, 500}, {1, 3, 20}, {3, 4, 10}},
			src:     1,
			dst:     4,
			k:       3,
		}, 30},
		{"", args{
			n:       5,
			flights: [][]int{{0, 1, 100}, {1, 4, 100}, {2, 4, 500}, {1, 3, 20}, {3, 4, 10}},
			src:     2,
			dst:     3,
			k:       3,
		}, -1},
		{"", args{
			n:       0,
			flights: [][]int{},
			src:     0,
			dst:     0,
			k:       0,
		}, -1},
		{"", args{
			n:       5,
			flights: [][]int{{1, 2, 10}, {2, 0, 7}, {1, 3, 8}, {4, 0, 10}, {3, 4, 2}, {4, 2, 10}, {0, 3, 3}, {3, 1, 6}, {2, 4, 5}},
			src:     0,
			dst:     4,
			k:       1,
		}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k); got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
