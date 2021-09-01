package leetcode

//有 n 个城市通过一些航班连接。给你一个数组flights ，其中flights[i] = [fromi, toi, pricei] ，表示该航班都从城市 fromi 开始，以价格 pricei 抵达 toi。
//
//现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k站中转的路线，使得从 src 到 dst 的 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/cheapest-flights-within-k-stops
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 解题思路(动态规划)
// 创建一个dp二维数组,长度为dp[k+2][n+1]
//
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	const intMax = int(^uint(0) >> 1)
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = intMax
		}
	}
	dp[0][src] = 0 //0到起点的初始值为0,为后面累加做准备
	for i := 1; i < k+2; i++ {
		for j := 1; j < n+1; j++ {
			//
			dp[i][j] = dp[i-1][j]
		}
		for j := 0; j < len(flights); j++ {
			flight := flights[j]
			from, to, price := flight[0], flight[1], flight[2]
			if dp[i-1][from] == intMax {
				continue
			}
			// dp[i-1][from]+price 为开始到上一站走过的路费+本站路费
			dp[i][to] = min(dp[i][to], dp[i-1][from]+price)
		}
	}
	if dp[k+1][dst] == intMax {
		return -1
	}
	return dp[k+1][dst]
}
