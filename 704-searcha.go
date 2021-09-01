package leetcode

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[((l+r)/2)] > target {
			r = (l+r)/2 - 1
		} else if nums[((l+r)/2)] < target {
			l = (l+r)/2 + 1
		} else {
			return (l + r) / 2
		}
	}
	return -1
}
