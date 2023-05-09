package main

import "fmt"

func romanToInt(s string) int {
	var result int
	var m map[byte]int = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i]] < m[s[i+1]] {
			result -= m[s[i]]
		} else {
			result += m[s[i]]
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
