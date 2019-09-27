package LeetCode

import "strings"

/*
	Problem #12 Integer to Roman https://leetcode.com/problems/integer-to-roman/
	16ms/3.3MB (Beats 18.89%/100.00%)
*/

func IntToRoman(num int) string {
	var romanNum strings.Builder

	if M := num / 1000; M > 0 {
		romanNum.WriteString(strings.Repeat("M", M))
		num %= 1000
	}

	if num >= 900 {
		romanNum.WriteString("CM")
		num -= 900
	}

	if num >= 500 {
		romanNum.WriteString("D")
		num -= 500
	}

	if num >= 400 {
		romanNum.WriteString("CD")
		num -= 400
	}

	if C := num / 100; C > 0 {
		romanNum.WriteString(strings.Repeat("C", C))
		num %= 100
	}

	if num >= 90 {
		romanNum.WriteString("XC")
		num -= 90
	}

	if num >= 50 {
		romanNum.WriteString("L")
		num -= 50
	}

	if num >= 40 {
		romanNum.WriteString("XL")
		num -= 40
	}

	if X := num / 10; X > 0 {
		romanNum.WriteString(strings.Repeat("X", X))
		num %= 10
	}

	if num >= 9 {
		romanNum.WriteString("IX")
		num -= 9
	}

	if num >= 5 {
		romanNum.WriteString("V")
		num -= 5
	}

	if num >= 4 {
		romanNum.WriteString("IV")
		num -= 4
	}

	if num > 0 {
		romanNum.WriteString(strings.Repeat("I", num))
	}

	return romanNum.String()
}
