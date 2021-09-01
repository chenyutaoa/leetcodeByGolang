package leetcode

//给定一个整数数组，判断是否存在重复元素。
//
//如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
func containsDuplicate(nums []int) bool {
	a := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := a[nums[i]]; ok {
			return true
		}
		a[nums[i]] = 1
	}
	return false
}
