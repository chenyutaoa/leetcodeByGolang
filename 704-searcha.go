package leetcode

//给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//func search(nums []int, target int) int {
//	l, r := 0, len(nums)-1
//	for l <= r {
//		if nums[((l+r)/2)] > target {
//			r = (l+r)/2 - 1
//		} else if nums[((l+r)/2)] < target {
//			l = (l+r)/2 + 1
//		} else {
//			return (l + r) / 2
//		}
//	}
//	return -1
//}

func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m -1
		} else {
			l = m +1
		}
	}
	return -1
}
