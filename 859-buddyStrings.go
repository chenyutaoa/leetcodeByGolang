package leetcode

//给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回true；否则返回 false 。
//
//交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。
//
//例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/buddy-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	idx := -1
	arr := []byte(s)
	m := make(map[byte]int, len(s))
	for i, _ := range arr {
		m[arr[i]]++
		if arr[i] != goal[i] {
			if idx == -1 {
				idx = i
			} else {
				arr[i], arr[idx] = s[idx], s[i]
				return string(arr) == goal
			}
		}
	}
	if idx != -1 {
		return false
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}
