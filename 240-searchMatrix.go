package leetcode

//func searchMatrix(matrix [][]int, target int) bool {
//	//最小值
//	min := func(a, b int) int {
//		if a > b {
//			return b
//		}
//		return a
//	}
//	//由于是升序数组，存储最右边索引
//	r := len(matrix[0]) - 1
//	//二分法获取是否存在目标值
//	find := func(arr []int) bool {
//		l := 0
//		for l <= r {
//			m := (l + r) / 2
//			if arr[m] == target {
//				return true
//			} else if arr[m] > target {
//				r = m - 1
//			} else {
//				l = m + 1
//			}
//		}
//		//设置最右边的索引，大值直接跳过
//		r = min(l, len(matrix[0])-1)
//		return false
//	}
//
//	for _, arr := range matrix {
//		if find(arr) {
//			return true
//		}
//	}
//	return false
//}

//coding标准答案
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	i := 0
	j := cols - 1
	for i < rows && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
