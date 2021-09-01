package leetcode

//给你两个非负整数 low 和 high 。请你返回 low 和 high 之间（包括二者）奇数的数目。

func countOdds(low int, high int) int {
	if (low & 1) == 0 { //偶数
		return (high - low + 1) / 2
	}
	return (high-low)/2 + 1
}
