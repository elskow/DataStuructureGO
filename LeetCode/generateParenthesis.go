package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(int, int, string)
	dfs = func(l, r int, s string) {
		if l == n && r == n {
			res = append(res, s)
			return
		}
		if l < n {
			dfs(l+1, r, s+"(")
		}
		if r < l {
			dfs(l, r+1, s+")")
		}
	}
	dfs(0, 0, "")
	return res
}

func main() {
	n := 3
	res := generateParenthesis(n)
	fmt.Println(strings.Join(res, "\n"))
}
