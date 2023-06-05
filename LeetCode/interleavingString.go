package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	memo := make(map[string]bool)

	return isInterleaveHelper(s1, s2, s3, memo)
}

func isInterleaveHelper(s1 string, s2 string, s3 string, memo map[string]bool) bool {
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}

	if len(s3) == 0 {
		return false
	}

	if result, ok := memo[s1+"#"+s2+"#"+s3]; ok {
		return result
	}

	result := false
	if len(s1) > 0 && s1[0] == s3[0] {
		result = result || isInterleaveHelper(s1[1:], s2, s3[1:], memo)
	}

	if len(s2) > 0 && s2[0] == s3[0] {
		result = result || isInterleaveHelper(s1, s2[1:], s3[1:], memo)
	}

	memo[s1+"#"+s2+"#"+s3] = result

	return result
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	println(isInterleave(s1, s2, s3))

	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbbaccc"
	println(isInterleave(s1, s2, s3))

	s1 = ""
	s2 = ""
	s3 = ""
	println(isInterleave(s1, s2, s3))
}
