package leetcode

func checkValidString(s string) bool {
	lf, rf := 0, 0
	for _, v := range s {
		if v == '(' {
			lf++
			rf++
		} else if v == ')' {
			lf = max(lf-1, 0)
			rf--
		} else {
			lf = max(lf-1, 0)
			rf++
		}
		if rf < 0 {
			return false
		}

	}
	return lf == 0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
