package LeetCode

/*
	Second attempt... Hopefully it will make sense this time
	Problem #2 Add Two Numbers https://leetcode.com/problems/add-two-numbers/
	8ms/5MB (Beats 89.25%/12.20%)
*/

func AddTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := false
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	var prev *ListNode

	head := result
	for true {
		if l1 != nil && l2 != nil {
			if carry {
				newVal := l1.Val + l2.Val + 1
				if newVal >= 10 {
					result.Val = newVal - 10
				} else {
					carry = false
					result.Val = newVal
				}
			} else {
				newVal := l1.Val + l2.Val
				if newVal >= 10 {
					result.Val = newVal - 10
					carry = true
				} else {
					result.Val = newVal
				}
			}
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			if carry {
				newVal := l1.Val + 1
				if newVal >= 10 {
					result.Val = newVal - 10
				} else {
					carry = false
					result.Val = newVal
				}
			} else {
				result.Val = l1.Val
			}
			l1 = l1.Next
		} else if l2 != nil {
			if carry {
				newVal := l2.Val + 1
				if newVal >= 10 {
					result.Val = newVal - 10
				} else {
					carry = false
					result.Val = newVal
				}
			} else {
				result.Val = l2.Val
			}
			l2 = l2.Next
		} else {
			if carry {
				result.Val = 1
			} else {
				prev.Next = nil
			}

			break
		}

		result.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		prev = result
		result = result.Next
	}

	return head
}
