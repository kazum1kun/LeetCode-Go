package main

import (
	"LeetCode"
	"fmt"
)

func mainATN() {
	l1 := &LeetCode.ListNode{
		Val: 4,
		Next: &LeetCode.ListNode{
			Val: 3,
			Next: &LeetCode.ListNode{
				Val: 1,
				Next: &LeetCode.ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	l2 := &LeetCode.ListNode{
		Val: 9,
		Next: &LeetCode.ListNode{
			Val:  2,
			Next: nil,
		},
	}

	ll1 := &LeetCode.ListNode{
		Val:  0,
		Next: nil,
	}
	ll2 := &LeetCode.ListNode{
		Val:  0,
		Next: nil,
	}

	printLinkedList(LeetCode.AddTwoNumbersV2(ll1, ll2))
	printLinkedList(LeetCode.AddTwoNumbersV2(l1, l2))
}

func printLinkedList(node *LeetCode.ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}
