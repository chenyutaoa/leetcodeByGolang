package leetcode

//给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。
//
//如果剩余字符少于 k 个，则将剩余字符全部反转。
//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

func reverseStr(s string, k int) string {
	b := []byte(s)
	reverse := func(b []byte) {
		l, r := 0, len(b)-1
		for l < r {
			b[l] ^= b[r]
			b[r] ^= b[l]
			b[l] ^= b[r]
			l++
			r--
		}
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 0; i < len(b)/(2*k)+1; i++ {
		a := b[min(2*k*i, len(b)):min(2*k*i+k, len(b))]
		if len(a) == 1 {
			continue
		}
		reverse(a)
	}
	return string(b)
}
