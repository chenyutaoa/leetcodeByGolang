package leetcode

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return nums[0]
	}
	res := nums[0]
	sum := 0
	for _, v := range nums {
		if sum > 0 {
			sum = v + sum
		} else {
			sum = v
		}
		if res < sum {
			res = sum
		}
	}
	return res
}
