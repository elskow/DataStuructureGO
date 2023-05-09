package main

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}
	var dict map[byte]string = map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var dfs func(int, string)
	dfs = func(index int, path string) {
		if index == len(digits) {
			result = append(result, path)
			return
		}
		for _, v := range dict[digits[index]] {
			dfs(index+1, path+string(v))
		}
	}
	dfs(0, "")
	return result
}

func main() {
	for _, v := range letterCombinations("23") {
		print(v, " ")
	}
}
