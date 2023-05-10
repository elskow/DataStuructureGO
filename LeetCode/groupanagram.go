package main

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	m := make(map[string][]string)
	for _, v := range strs {
		key := sortString(v)
		m[key] = append(m[key], v)
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func sortString(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		for j := i; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return string(b)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for _, v := range groupAnagrams(strs) {
		for _, vv := range v {
			print(vv, " ")
		}
		println()
	}
}
