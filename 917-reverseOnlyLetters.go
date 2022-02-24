package leetcode

func reverseOnlyLetters(s string) string {
	checkLetter := func(i uint8) bool {
		return (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z')
	}
	result := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		left, right := checkLetter(result[l]), checkLetter(result[r])
		if left && right {
			result[l], result[r] = result[r], result[l]
			l++
			r--
			continue
		}
		if !left {
			l++
		}
		if !right {
			r--
		}
	}
	return string(result)
}
