package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var maxLen int
	var start int
	var charMap [256]int
	for i := 0; i < len(s); i++ {
		if charMap[s[i]] > start {
			start = charMap[s[i]]
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		charMap[s[i]] = i + 1
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
