package LeetCode

import (
	"math/big"
	"strconv"
)

/*
	VERY INEFFICIENT AND BAD IMPL!!! Maybe I shouldn't code while asleep...
	Problem #2 Add Two Numbers https://leetcode.com/problems/add-two-numbers/
	24ms/5.7MB (Beats 9.06%/7.32%)
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1Str, num2Str := "", ""
	for true {
		if l1 != nil && l2 != nil {
			num1Str = strconv.Itoa(l1.Val) + num1Str
			num2Str = strconv.Itoa(l2.Val) + num2Str
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			num1Str = strconv.Itoa(l1.Val) + num1Str
			l1 = l1.Next
		} else if l2 != nil {
			num2Str = strconv.Itoa(l2.Val) + num2Str
			l2 = l2.Next
		} else {
			break
		}
	}

	// Boy, some test cases use 30+ digit numbers...
	num1, num2 := new(big.Int), new(big.Int)
	num1.SetString(num1Str, 10)
	num2.SetString(num2Str, 10)
	sum := new(big.Int)
	sum.Add(num1, num2)

	sumStr := sum.String()
	val, _ := strconv.Atoi(string(sumStr[len(sumStr)-1]))
	reversedList := &ListNode{
		Val:  val,
		Next: nil,
	}

	head := reversedList
	for i := len(sumStr) - 2; i >= 0; i-- {
		val, _ := strconv.Atoi(string(sumStr[i]))
		reversedList.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		reversedList = reversedList.Next
	}

	return head
}
