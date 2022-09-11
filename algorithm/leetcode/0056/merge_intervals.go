package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	var cs []int
	for _, interval := range intervals {
		if len(res) == 0 || cs[1] < interval[0] {
			res = append(res, interval)
			cs = interval
		} else if cs[1] < interval[1] {
			cs[1] = interval[1]
		}
	}
	return res
}
