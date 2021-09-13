package leetcode

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		m := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			m[x*x+y*y] ++
		}
		for _, v := range m {
			ans += v * (v - 1)
		}
	}
	return ans
}
