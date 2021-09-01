package leetcode

//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
//元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
func reverseVowels(s string) string {
	m := map[uint8]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	b := []byte(s)
	var a []uint8
	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := m[s[i]]; ok {
			a = append(a, s[i])
		}
	}
	idx := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			b[i] = a[idx]
			idx++
		}
	}
	return string(b)
}

func reverseVowels1(s string) string {
	if len(s) <= 1 {
		return s
	}
	exist := func(c uint8) bool {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			return true
		}
		return false
	}
	b := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !exist(s[l]) {
			l++
		}
		for l < r && !exist(s[r]) {
			r--
		}
		if b[l] != b[r] {
			b[l] ^= b[r]
			b[r] ^= b[l]
			b[l] ^= b[r]
		}
		l++
		r--
	}
	return string(b)
}
