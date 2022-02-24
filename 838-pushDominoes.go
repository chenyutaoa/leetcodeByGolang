package leetcode

//n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
//
//每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
//
//如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
//
//就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。
//
//给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：
//
//dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
//dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
//dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。
//返回表示最终状态的字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/push-dominoes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func pushDominoes(dominoes string) string {
	var left int
	for i, v := range dominoes {
		if v == '.' {
			if i == len(dominoes)-1 {
				return pushing(dominoes, left, i)
			}
			continue
		}
		dominoes = pushing(dominoes, left, i)
		left = i
	}
	return dominoes
}
func pushing(dominoes string, l, r int) string {
	s := []byte(dominoes)
	if dominoes[l] == 'R' && dominoes[r] == 'L' && r-l > 2 { //方向相对
		for l < r && r-l > 2 {
			s[l+1] = 'R'
			s[r-1] = 'L'
			r--
			l++
		}
		return string(s)
	}
	if dominoes[l] == dominoes[r] { //方向相同
		for i := l; i <= r; i++ {
			s[i] = dominoes[l]
		}
		return string(s)
	}
	if dominoes[l] == '.' && dominoes[r] == 'L' { //方向相同
		for i := l; i < r; i++ {
			s[i] = dominoes[r]
		}
		return string(s)
	}
	if dominoes[r] == '.' && dominoes[l] == 'R' { //方向相同
		for i := l; i < r; i++ {
			s[i+1] = dominoes[l]
		}
		return string(s)
	}
	return string(s)
}

func pushDominoes1(dominoes string) string {
	s := []byte(dominoes)
	i, n, left := 0, len(s), byte('L')
	for i < n {
		j := i
		for j < n && s[j] == '.' { // 找到一段连续的没有被推动的骨牌
			j++
		}
		right := byte('R')
		if j < n {
			right = s[j]
		}
		if left == right { // 方向相同，那么这些竖立骨牌也会倒向同一方向
			for i < j {
				s[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' { // 方向相对，那么就从两侧向中间倒
			k := j - 1
			for i < k {
				s[i] = 'R'
				s[k] = 'L'
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	return string(s)
}
