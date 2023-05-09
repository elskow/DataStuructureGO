package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	s := "abc"
	t := "ahbgdc"
	print(isSubsequence(s, t))
}
