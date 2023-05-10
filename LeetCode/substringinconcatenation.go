package main

func findSubstring(s string, words []string) []int {
	var res []int
	if len(words) == 0 {
		return res
	}
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	n := len(words[0])
	for i := 0; i < len(s)-n*len(words)+1; i++ {
		seen := make(map[string]int)
		j := 0
		for ; j < len(words); j++ {
			w := s[i+j*n : i+(j+1)*n]
			if _, ok := m[w]; !ok {
				break
			}
			seen[w]++
			if seen[w] > m[w] {
				break
			}
		}
		if j == len(words) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	for _, v := range findSubstring(s, words) {
		print(v, " ")
	}
	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "word"}
	for _, v := range findSubstring(s, words) {
		print(v, " ")
	}
}
