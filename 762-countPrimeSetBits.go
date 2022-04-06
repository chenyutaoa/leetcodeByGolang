package leetcode

//给你两个整数left和right ，在闭区间 [left, right]范围内，统计并返回 计算置位位数为质数 的整数个数。
//
//计算置位位数 就是二进制表示中 1 的个数。
//
//例如， 21的二进制表示10101有 3 个计算置位。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func countPrimeSetBits(left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		cur := i
		sumOne := 0
		for cur > 0 {
			if cur&1 == 1 {
				sumOne++
			}
			cur = cur >> 1
		}
		if isPrime(sumOne){
			count++
		}
	}
	return count
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}