package main

import (
	"github.com/larryrun/leetcode-go/utils"
)

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
 */

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	startIdx := 0
	existing := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		dupIdx, ok := existing[s[i]]
		if ok {
			if dupIdx >= startIdx {
				startIdx = dupIdx + 1
			}
		}
		existing[s[i]] = i
		lenSoFar := (i - startIdx) + 1
		if lenSoFar > maxLen {
			maxLen = lenSoFar
		}
	}
	return maxLen
}

func main() {
	utils.AssertEq(3, lengthOfLongestSubstring("abcabcbb"))
	utils.AssertEq(1, lengthOfLongestSubstring("bbbbb"))
	utils.AssertEq(3, lengthOfLongestSubstring("pwwkew"))
	utils.AssertEq(2, lengthOfLongestSubstring("au"))
	utils.AssertEq(1, lengthOfLongestSubstring(" "))
	utils.AssertEq(2, lengthOfLongestSubstring("abba"))
}