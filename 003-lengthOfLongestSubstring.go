package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res, low := 0, 0
	m := make(map[int32]int)
	for high, b := range s {
		if _, ok := m[b]; ok {
			low = max(m[b]+1,low)
		}
		res = max(res, high-low+1)
		m[b] = high
	}
	return res
}
