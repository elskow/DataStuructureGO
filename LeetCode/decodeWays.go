package main

func numDecodings(s string) int {
	memo := make(map[string]int)

	return numDecodingsHelper(s, memo)
}

func numDecodingsHelper(s string, memo map[string]int) int {
	if len(s) == 0 {
		return 1
	}

	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	if result, ok := memo[s]; ok {
		return result
	}

	result := numDecodingsHelper(s[1:], memo)
	if s[0] == '1' || (s[0] == '2' && s[1] <= '6') {
		result += numDecodingsHelper(s[2:], memo)
	}

	memo[s] = result

	return result
}

func main() {
	s := "12"
	println(numDecodings(s))
	s = "226"
	println(numDecodings(s))
}
