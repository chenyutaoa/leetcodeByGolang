package leetcode

func chalkReplacer(chalk []int, k int) int {
	if len(chalk) == 1 {
		return 0
	}
	sum := 0
	for i := 0; i < len(chalk); i++ {
		sum += chalk[i]
		if sum > k {
			return i
		}
		chalk[i] = sum
	}
	s := k % sum
	l, r := 0, len(chalk)-1
	for l <= r {
		m := (l + r) / 2
		if chalk[m] == s {
			return m + 1
		} else if chalk[m] < s {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
