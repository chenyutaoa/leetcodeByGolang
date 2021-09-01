package leetcode

//给你一个整数 n 。按下述规则生成一个长度为 n + 1 的数组 nums ：
//
//nums[0] = 0
//nums[1] = 1
//当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
//当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]
//返回生成数组 nums 中的 最大 值。

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}

	m := []uint8{0, 1}
	r := uint8(0)
	for i := 2; i <= n; i++ {
		if (i & 1) == 0 { //偶数
			m = append(m, m[i/2])
		} else {
			m = append(m, m[i/2]+m[i/2+1])
		}
		if r < m[i] {
			r = m[i]
		}
	}
	return int(r)
}
