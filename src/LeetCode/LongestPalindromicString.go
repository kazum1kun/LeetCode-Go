package LeetCode

/*
	Problem #5 Two Sum https://leetcode.com/problems/longest-palindromic-substring/
	Basic idea is that iterate through each character, and expand both ways to
	see if palindrome can be found
*/

func LongestPalindrome(s string) string {
	currentBest, currentBestLen := "", 0

	for i := 1; i < len(s)-1; i++ {
		sameChars := true
		for diff := 1; i-diff >= 0 && i+diff < len(s); diff++ {
			if s[i-diff] == s[i+diff] {
				if sameChars && s[i-diff] != s[i] {
					sameChars = false
				}
				if currentBestLen < diff*2+1 {
					currentBest = s[i-diff : i+diff+1]
					currentBestLen = len(currentBest)
				}
				continue
			}
		}
	}

	return currentBest
}
