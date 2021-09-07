package leetcode

func balancedStringSplit(s string) int {
	if len(s)%2 != 0 {
		return -1
	}
	if len(s) <= 2 {
		return len(s)
	}
	res := 0
	l, r := 0, 0
	for i := 0; i < len(s); i += 2 {
		if l == 0 && r == 0 && s[i] != s[i+1] {
			res++
			continue
		}
		if s[i+1] == 'L' {
			l++
		}else {
			r++
		}
		if s[i] == 'L' {
			l++
		}else {
			r++
		}
		if l == r {
			res++
			l, r = 0, 0
			continue
		}
	}
	return res
}
