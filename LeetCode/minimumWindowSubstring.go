package main

func minWindow(s string, t string) string {
	var res string
	var left, right, minLen, count int = 0, 0, len(s) + 1, 0
	var m = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	for right < len(s) {
		if m[s[right]] > 0 {
			count++
		}
		m[s[right]]--
		right++
		for count == len(t) {
			if right-left < minLen {
				minLen = right - left
				res = s[left:right]
			}
			m[s[left]]++
			if m[s[left]] > 0 {
				count--
			}
			left++
		}
	}
	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	println(minWindow(s, t))
	s = "a"
	t = "a"
	println(minWindow(s, t))
}
