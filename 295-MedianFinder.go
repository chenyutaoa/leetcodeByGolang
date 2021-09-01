package leetcode

//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
//例如，
//
//[2,3,4] 的中位数是 3
//
//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
//设计一个支持以下两种操作的数据结构：
//
//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
//double findMedian() - 返回目前所有元素的中位数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-median-from-data-stream
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MedianFinder struct {
	Nums []int//有序数组
}

/** initialize your data structure here. */
func MedianFinderConstructor() MedianFinder {
	return MedianFinder{
		Nums: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.Nums) == 0 {
		this.Nums = append(this.Nums, num)
		return
	}
	i := this.insertIdx(num)//获取新插入数应该存放的索引
	this.Nums = append(this.Nums, 0)     // 切片扩展1个空间
	copy(this.Nums[i+1:], this.Nums[i:]) // a[i:]向后移动1个位置
	this.Nums[i] = num
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Nums) == 0 {
		return 0
	}
	if len(this.Nums) == 1 {
		return float64(this.Nums[0])
	}
	if len(this.Nums)&1 == 0 { //偶数
		return float64(this.Nums[len(this.Nums)/2-1] + this.Nums[len(this.Nums)/2]) / 2
	} else {
		return float64(this.Nums[len(this.Nums)/2])
	}
}

//获取新插入数应该存放的索引
func (this *MedianFinder) insertIdx(num int) int {
	if this.Nums[0] >= num {
		return 0
	}
	if this.Nums[len(this.Nums)-1] <= num {
		return len(this.Nums)
	}
	l, r := 0, len(this.Nums)-1
	for l <= r {
		m := (l + r) / 2
		if this.Nums[m] == num {
			return m
		} else if this.Nums[m] > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := MedianFinderConstructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
