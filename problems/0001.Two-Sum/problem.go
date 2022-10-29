package leetcode

func twoSum(nums []int, target int) []int {
	complementaryMap := make(map[int]int, 0)

	for i, value := range nums {
		if j, ok := complementaryMap[target-value]; ok {
			return []int{j, i}
		}
		complementaryMap[value] = i
	}
	return nil
}
