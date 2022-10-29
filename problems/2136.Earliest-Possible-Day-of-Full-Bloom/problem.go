package leetcode

import "sort"

func earliestFullBloom(plantTime []int, growTime []int) int {
	indices := make([]int, 0, len(growTime))
	for i := 0; i < len(growTime); i++ {
		indices = append(indices, i)
	}

	sort.Slice(indices, func(i, j int) bool {
		return growTime[indices[i]] > growTime[indices[j]]
	})

	res, currentDay := 0, 0
	for _, i := range indices {
		currentDay += plantTime[i]
		// case when planting and growing a seed takes less time than growing some previous seed
		res = max(res, currentDay+growTime[i])
	}

	return res
}

func max(n, m int) int {
	if n < m {
		return m
	}
	return n
}
