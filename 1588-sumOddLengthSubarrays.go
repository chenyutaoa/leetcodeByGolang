package leetcode

func sumOddLengthSubarrays(arr []int) int {
	s:=0
	for i,v:=range arr{
		ss := v
		s += v
		ii:=1
		if i == len(arr)-2{
			continue
		}
		for j := i+1;j<len(arr);j++{
			ii++
			ss += arr[j]
			if ii %2==0{
				continue
			}
			s += ss
		}
	}
	return s
}
