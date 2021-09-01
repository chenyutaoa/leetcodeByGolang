package leetcode

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
// 
//
//进阶：
//
//尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
//你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func rotate(nums []int, k int) []int {
	if k%len(nums) == 0 {
		return nums
	}
	for i := 0; i < len(nums)/2+1; i++ {
		//nums[(i+k)%len(nums)] ^= nums[i]
		//nums[i] ^= nums[(i+k)%len(nums)]
		//nums[(i+k)%len(nums)] ^= nums[i]
	}
	return nums
}
