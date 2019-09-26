package LeetCode

/*
	Problem #5 Two Sum https://leetcode.com/problems/longest-palindromic-substring/
	Basic idea is that iterate through each character, and expand both ways to
	see if palindrome can be found
	12ms/2.2MB (Beats 57.83%/100.00%(!!!))
*/

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	currentBest, currentBestLen := s[0:1], 1
	sameChars := true

	if len(s) >= 2 {
		if s[0] == s[1] {
			currentBest = s[0:2]
			currentBestLen = 2
		}
	}

	for i := 1; i < len(s)-1; i++ {
		oddDiff, evenDiff := 0, 0
		for diff := 1; i-diff >= 0 && i+diff < len(s); diff++ {
			operation := false
			// Odd length palindrome
			if s[i-diff] == s[i+diff] && diff == oddDiff+1 {
				if sameChars && s[i-diff] != s[i] {
					sameChars = false
				}
				if currentBestLen < diff*2+1 {
					if sameChars && i+diff+1 < len(s) && s[i+diff+1] == s[i] {
						currentBest = s[i-diff : i+diff+2]
					} else {
						currentBest = s[i-diff : i+diff+1]
					}
					currentBestLen = len(currentBest)
				}
				operation = true
				oddDiff++
			}
			// Even length palindrome
			if s[i] == s[i+1] && i+diff+1 < len(s) && s[i-diff] == s[i+diff+1] && diff == evenDiff+1 {
				if currentBestLen < diff*2+2 {
					currentBest = s[i-diff : i+diff+2]
					currentBestLen = len(currentBest)
				}
				operation = true
				evenDiff++
			}
			if currentBestLen < 2 && i+1 < len(s) && s[i+1] == s[i] {
				currentBest = s[i : i+2]
				currentBestLen = 2
				break
			}
			if !operation {
				break
			}
		}
	}

	return currentBest
}
