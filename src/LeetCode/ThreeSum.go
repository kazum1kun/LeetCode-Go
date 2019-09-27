package LeetCode

import "sort"

/*
	Problem #15 3Sum https://leetcode.com/problems/3sum/
	Basically iterate through each int and attempt 2Sum for the given integer
	2440ms/272.5MB (Beat 5.09%/50.00%)
	Optimization can be used, but that would be the work for another time...
*/

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	numsMap := make(map[int]int)

	flag := true
	for k, v := range nums {
		if v != 0 {
			flag = false
		}
		numsMap[v] = k
	}

	if flag {
		return [][]int{{0, 0, 0}}
	}

	solution := make([][]int, 0)

	for i, e := range nums {
		for i2, e2 := range nums {
			i3, prs := numsMap[0-e-e2]
			if prs && i != i2 && i2 != i3 && i != i3 {
				newSolution := []int{e, e2, 0 - e - e2}
				sort.Ints(newSolution)
				add := true
				for ai := 0; ai < len(solution); ai++ {
					if newSolution[0] == solution[ai][0] &&
						newSolution[1] == solution[ai][1] &&
						newSolution[2] == solution[ai][2] {
						add = false
						break
					}
				}
				if add {
					solution = append(solution, newSolution)
				}
			}
		}
	}

	return solution
}
