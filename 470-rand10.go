package leetcode

import "math/rand"

func rand10() int {
	f, s := 10, 10
	for f > 6 {
		f = rand7()
	}
	for s > 5 {
		s = rand7()
	}
	if (f & 1) == 0 {
		return s
	}
	return 5 + s
}

func rand7() int {
	return rand.Intn(6) + 1
}
