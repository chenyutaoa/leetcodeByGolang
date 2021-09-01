package leetcode

// 暴力遍历
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			continue
		}
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// map集合
func twoSum(nums []int, target int) []int {
	var m map[int]int /*创建集合 */
	m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok && v != i {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
