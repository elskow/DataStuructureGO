package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m1, m2 [256]int
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	print(isIsomorphic(s, t))
}
