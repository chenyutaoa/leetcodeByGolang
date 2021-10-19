package leetcode

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	len1, len2 := len(nums1), len(nums2)
//	if len1 == 0 {
//		return median(nums2)
//	}
//	if len2 == 0 {
//		return median(nums1)
//	}
//	if nums1[len1-1] <= nums2[0] {
//		return median(append(nums1, nums2...))
//	}
//	if nums2[len2-1] <= nums1[0] {
//		return median(append(nums2, nums1...))
//	}
//	var findMedian func(num []int, val int) int
//	findMedian = func(nums []int, val int) int {
//		l, r := 0, len(nums)-1
//		if val >= nums[r]{
//			return r+1
//		}
//		if val <= nums[0]{
//			return 0
//		}
//		for l <= r {
//			m := (l + r) / 2
//			if nums[m] == val {
//				return m
//			} else if nums[m] > val {
//				r = m - 1
//			} else {
//				l = m + 1
//			}
//		}
//		return l
//	}
//	for i := 0; i < len2; i++ {
//		idx := findMedian(nums1, nums2[i])
//		nums1 = append(nums1,0)
//		copy(nums1[idx+1:], nums1[idx:]) // a[i:]向后移动1个位置
//		nums1[idx] = nums2[i]
//	}
//	return median(nums1)
//}
//
//func median(num []int) float64 {
//	if len(num) == 0 {
//		return 0
//	}
//	if len(num) == 1 {
//		return float64(num[0])
//	}
//	if len(num)%2 == 0 {
//		return (float64(num[len(num)/2]) + float64(num[len(num)/2-1])) / 2
//	}
//	return float64(num[len(num)/2])
//}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

type littleHeap struct {
	dataArray []int
}

func (lh *littleHeap) String() {
	for _, i := range lh.dataArray {
		print(i)
	}
}

func (lh *littleHeap) insert(data int) {
	lh.dataArray = append(lh.dataArray, data)
	lh.rebuildLittleHeapAfterAdd()
}

func (lh *littleHeap) deleteHeader() {
	if len(lh.dataArray) <= 0 {
		return
	}
	value := lh.dataArray[len(lh.dataArray)-1]
	lh.dataArray = lh.dataArray[:len(lh.dataArray)-2]
	lh.dataArray[0] = value
	lh.rebuildAfterDelete()
}

func (lh *littleHeap) replaceHeader(data int) {
	if len(lh.dataArray) > 0 {
		lh.dataArray[0] = data
		lh.rebuildAfterDelete()
	}
}

func (lh *littleHeap) rebuildLittleHeapAfterAdd() {
	idx := len(lh.dataArray) - 1
	for {
		fatherIdx := lh.getFatherNodeIdx(idx)
		currentData := lh.dataArray[idx]
		fatherData := lh.dataArray[fatherIdx]
		if currentData < fatherData {
			lh.dataArray[idx], lh.dataArray[fatherIdx] = lh.dataArray[fatherIdx], lh.dataArray[idx]
			idx = fatherIdx
		} else {
			break
		}
	}
}

func (lh *littleHeap) rebuildAfterDelete() {
	idx := 0
	maxIdx := len(lh.dataArray) - 1
	for {
		leftIdx := lh.getLeftNodeIdx(idx)
		rightIdx := lh.getRightNodeIdx(idx)

		var leftValue, rightValue int
		if leftIdx <= maxIdx {
			leftValue = lh.dataArray[leftIdx]
		} else {
			leftValue = 65535
		}

		if rightIdx <= maxIdx {
			rightValue = lh.dataArray[rightIdx]
		} else {
			rightValue = 65535
		}

		if leftValue < rightValue {
			if lh.dataArray[idx] < leftValue {
				break
			}

			lh.dataArray[leftIdx], lh.dataArray[idx] = lh.dataArray[idx], lh.dataArray[leftIdx]
			idx = leftIdx
		} else {
			if lh.dataArray[idx] < rightValue {
				break
			}

			lh.dataArray[rightIdx], lh.dataArray[idx] = lh.dataArray[idx], lh.dataArray[rightIdx]
			idx = rightIdx
		}
	}
}

func (lh *littleHeap) getLeftNodeIdx(currentIdx int) int {
	if currentIdx == 0 {
		return 1
	} else {
		return currentIdx * 2
	}
}

func (lh *littleHeap) getRightNodeIdx(currentIdx int) int {
	if currentIdx == 0 {
		return 2
	} else {
		return currentIdx*2 + 1
	}
}

func (lh *littleHeap) getFatherNodeIdx(currentIdx int) int {
	return currentIdx / 2
}

func (lh *littleHeap) getTop2Value() (int, int) {
	if len(lh.dataArray) == 1 {
		return lh.dataArray[0], lh.dataArray[0]
	} else if len(lh.dataArray) == 2 {
		return lh.dataArray[0], lh.dataArray[1]
	}

	li := lh.getLeftNodeIdx(0)
	ri := lh.getRightNodeIdx(0)
	min := min(lh.dataArray[li], lh.dataArray[ri])
	return lh.dataArray[0], int(min)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num1len := len(nums1)
	num2len := len(nums2)

	sumlen := num1len + num2len
	lh := new(littleHeap)

	var shortnums, longnums []int
	if num1len < num2len {
		shortnums = nums1
		longnums = nums2
	} else {
		shortnums = nums2
		longnums = nums1
	}

	for _, shortnum := range shortnums {
		lh.insert(shortnum)
	}

	var curLongnumsIdx int
	idx := (sumlen + 2) / 2
	remain := idx - len(shortnums)
	for i := 0; i < remain; i++ {
		lh.insert(longnums[i])
	}
	curLongnumsIdx = remain
	for i := curLongnumsIdx; i < len(longnums); i++ {
		if lh.dataArray[0] < longnums[i] {
			lh.replaceHeader(longnums[i])
		}
	}

	var v1, v2 int
	if (sumlen % 2) == 0 {
		v1, v2 = lh.getTop2Value()
	} else {
		v1 = lh.dataArray[0]
		v2 = v1
	}

	return float64((v1 + v2)) / 2
}
