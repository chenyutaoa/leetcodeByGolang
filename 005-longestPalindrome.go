package leetcode

//给你一个字符串 s，找到 s 中最长的回文子串。

//func longestPalindrome(s string) string {
//	m := make(map[string]bool)
//	res := string(s[0])
//	for idx, c := range s {
//		a := string(c)
//		for _, cc := range s[idx+1:] {
//			a += string(cc)
//			if _, ok := m[a]; ok {
//				continue
//			}
//			is := isPalindrome(a)
//			m[a] = is
//			if is && len(res) < len(a) {
//				res = a
//			}
//		}
//	}
//	return res
//}
//
//func isPalindrome(s string) bool {
//	l, r := 0, len(s)-1
//	for l < r {
//		if s[l] != s[r] {
//			return false
//		}
//		l++
//		r--
//	}
//	return true
//}

func longestPalindrome(s string) string {
	l, r := 0, 0
	for idx, _ := range s {
		//奇数回文
		l, r = getPalindrome(s, l, r, idx, idx)
		//偶数回文
		l, r = getPalindrome(s, l, r, idx, idx+1)
	}
	return s[l : r+1]
}

func getPalindrome(s string, l, r, ll, rr int) (int, int) {
	for ; ll >= 0 && rr < len(s) && s[ll] == s[rr]; ll, rr = ll-1, rr+1 {

	}
	ll, rr = ll+1, rr-1
	if rr-ll > r-l {
		return ll, rr
	}
	return l, r
}
