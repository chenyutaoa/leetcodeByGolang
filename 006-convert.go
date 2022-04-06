package leetcode

//
//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行Z 字形排列。
//
//比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zigzag-conversion
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	max := 2 * (numRows - 1)
	res := ""
	for i := 0; i < numRows; i++ {
		idx := i
		if i == numRows-1 || i == 0 {
			for idx < len(s) {
				res += string(s[idx])
				idx += max
			}
		} else {
			ii := 0
			for idx < len(s) {
				res += string(s[idx])
				if ii%2 == 0 {
					idx += max - i*2
				} else {
					idx += max - (max - i*2)
				}
				ii++
			}

		}

	}
	return res
}

//func convert(s string, numRows int) string {
//	if numRows == 1 {
//		return s
//	}
//	rows := make([]string, numRows)
//	n := 2*numRows - 2
//	for i, char := range s {
//		x := i % n
//		rows[min(x, n-x)] += string(char)
//	}
//	return strings.Join(rows, "")
//}
