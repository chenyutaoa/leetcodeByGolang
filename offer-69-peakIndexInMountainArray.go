package leetcode

import "fmt"

//
//符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：
//
//arr.length >= 3
//存在 i（0 < i< arr.length - 1）使得：
//arr[0] < arr[1] < ... arr[i-1] < arr[i]
//arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//给定由整数组成的山峰数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i，即山峰顶部。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/B1IidL
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var c1 = make(chan int)
var c2 = make(chan int)

func peakIndexInMountainArray(arr []int) int {
	//if len(arr) < 3 {
	//	return -1
	//}
	//l, r := 0, len(arr)-1
	//for l < r {
	//	m := (l + r) / 2
	//	lm, rm := (m+l)/2, (m+r)/2
	//	mv, lmv, rmv := arr[m], arr[lm], arr[rm]
	//	if l+1 == m && r-1 == m {
	//		return m
	//	} else if lmv > mv {
	//		r = m
	//	} else if rmv > mv {
	//		l = m
	//	}else if mv >=lmv && mv >= rmv{
	//		r = rm
	//		l = lm
	//	}
	//}
	c := &C1{}
	fmt.Sprintf("s : %v", c)
	return -1
}

type C1 struct {
	De string
}
