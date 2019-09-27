package main

import (
	"LeetCode"
	"fmt"
)

func twoSumMain() {
	//fmt.Println(LeetCode.TwoSum([]int{3,2,4}, 6))
	l1 := &LeetCode.ListNode{
		Val: 5,
		Next: &LeetCode.ListNode{
			Val: 2,
			Next: &LeetCode.ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	l2 := &LeetCode.ListNode{
		Val: 4,
		Next: &LeetCode.ListNode{
			Val: 0,
			Next: &LeetCode.ListNode{
				Val: 0,
				Next: &LeetCode.ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}
	ret := LeetCode.AddTwoNumbers(l1, l2)

	for ret != nil {
		fmt.Print(ret.Val)
		ret = ret.Next
	}
}
