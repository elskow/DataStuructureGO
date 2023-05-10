package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	memo := make(map[string]bool)
	return solve(s, p, memo)
}

func solve(s string, p string, memo map[string]bool) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s) == 0 {
		return allStars(p)
	}
	if val, ok := memo[s+" "+p]; ok {
		return val
	}
	result := false
	if p[0] == '*' {
		result = solve(s[1:], p, memo) || solve(s, p[1:], memo)
	} else {
		result = match(s[0], p[0]) && solve(s[1:], p[1:], memo)
	}
	memo[s+" "+p] = result
	return result
}

func match(s byte, p byte) bool {
	return s == p || p == '?'
}

func allStars(p string) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isMatch("aa", "a"))        // false
	fmt.Println(isMatch("aa", "*"))        // true
	fmt.Println(isMatch("cb", "?a"))       // false
	fmt.Println(isMatch("adceb", "*a*b"))  // true
	fmt.Println(isMatch("acdcb", "a*c?b")) // false
}
