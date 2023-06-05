package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func isScramble(s1 string, s2 string) bool {
	memo := make(map[string]bool)
	return isScrambleHelper(s1, s2, memo)
}

func isScrambleHelper(s1 string, s2 string, memo map[string]bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	if s1 == s2 {
		return true
	}
	if memo[s1+" "+s2] {
		return true
	}
	if memo[s2+" "+s1] {
		return true
	}
	if !isAnagram(s1, s2) {
		return false
	}
	for i := 1; i < len(s1); i++ {
		if isScrambleHelper(s1[:i], s2[:i], memo) && isScrambleHelper(s1[i:], s2[i:], memo) {
			memo[s1+" "+s2] = true
			return true
		}
		if isScrambleHelper(s1[:i], s2[len(s2)-i:], memo) && isScrambleHelper(s1[i:], s2[:len(s2)-i], memo) {
			memo[s1+" "+s2] = true
			return true
		}
	}
	return false
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	counts := make(map[rune]int)
	for _, c := range s1 {
		counts[c]++
	}
	for _, c := range s2 {
		counts[c]--
	}
	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "great"
	s2 := "rgeat"
	println(isScramble(s1, s2))
}
