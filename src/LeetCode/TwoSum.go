package LeetCode

/*
	Problem #1 Two Sum https://leetcode.com/problems/two-sum/
	4ms/3.8MB (Beats 95.15%/13.46%)
	An obvious O(n^2) solution exists (and prolly trade off runtime for memory)
*/
func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for k, v := range nums {
		numsMap[v] = k
	}

	for i, e := range nums {
		i2, prs := numsMap[target-e]
		if prs && i != i2 {
			return []int{i, i2}
		}
	}

	return nil
}
