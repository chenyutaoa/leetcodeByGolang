package leetcode

import "sort"

//func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
//	m := make(map[int]int)
//	for k > 0 {
//		mv, idx := 0, 0
//		for i, v := range profits {
//			if _, ok := m[i]; ok {
//				continue
//			}
//			if capital[i] <= w && v >= mv {
//				idx = i
//				mv = v
//			}
//		}
//		w += mv
//		m[idx] = 0
//		k--
//	}
//	return w
//}

type pair struct {
	c, p int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	pairs := make([]pair, len(profits))
	for i, v := range profits {
		pairs[i] = pair{
			c: capital[i],
			p: v,
		}
	}
	//排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].c < pairs[j].c
	})
	m := make(map[int]int)
	for k > 0 {
		mv, idx := 0, 0
		for i, v := range pairs {
			if _, ok := m[i]; ok {
				continue
			}
			if capital[i] <= w && v.p >= mv {
				idx = i
				mv = v.p
			}
		}
		w += mv
		m[idx] = 0
		k--
	}
	return w
}
