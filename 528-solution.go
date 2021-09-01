package leetcode

import "math/rand"

type Solution struct {
	w      []int
	length int
}

func SolutionConstructor(w []int) Solution {
	var s Solution
	s.w = w
	for _, v := range w {
		s.length += v
	}
	return s
}

func (this *Solution) PickIndex() int {
	if len(this.w) == 0 {
		return -1
	}
	if len(this.w) == 1 {
		return 0
	}
	ra := rand.Intn(this.length)
	sum := 0
	for i, v := range this.w {
		sum += v
		if sum > ra {
			return i
		}
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := MedianFinderConstructor(w);
 * param_1 := obj.PickIndex();
 */
