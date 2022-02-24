package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n2len := len(nums2)
	map1 := make(map[int]int, n2len)
	for i, v := range nums2 {
		map1[v] = i
	}
	for i, v := range nums1 {
		if map1[v] != n2len-1 && nums2[map1[v]+1] > v {
			nums1[i] = nums2[map1[v]+1]
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}
