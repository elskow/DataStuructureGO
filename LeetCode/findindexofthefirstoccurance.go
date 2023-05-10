package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var (
		i, j int
	)
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
		if j == len(needle) {
			return i - j
		}
	}
	return -1
}

func main() {
	println(strStr("hello", "ll"))
	println(strStr("aaaaa", "bba"))
	println(strStr("aaaaa", ""))
	println(strStr("aaaaa", "a"))
	println(strStr("mississippi", "issip"))
}
