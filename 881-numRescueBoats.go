package leetcode

import "sort"

//第i个人的体重为people[i]，每艘船可以承载的最大重量为limit。
//
//每艘船最多可同时载两人，但条件是这些人的重量之和最多为limit。
//
//返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/boats-to-save-people
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func numRescueBoats(people []int, limit int) int {
	length := 0
	sort.Ints(people)
	l, r := 0, len(people)-1
	for l <= r {
		if people[r]+people[l] <= limit {
			l++
			continue
		}
		r--
		length++
	}
	return length
}
