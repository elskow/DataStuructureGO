package main

func combine(n int, k int) [][]int {
	var res [][]int
	var tmp []int
	dfs(n, k, 1, &tmp, &res)
	return res
}

func dfs(n, k, start int, tmp *[]int, res *[][]int) {
	if len(*tmp) == k {
		var t []int
		t = append(t, *tmp...)
		*res = append(*res, t)
		return
	}
	for i := start; i <= n; i++ {
		*tmp = append(*tmp, i)
		dfs(n, k, i+1, tmp, res)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}

func main() {
	n := 4
	k := 2
	for _, v := range combine(n, k) {
		for _, v1 := range v {
			print(v1, " ")
		}
		println()
	}

	n = 1
	k = 1
	for _, v := range combine(n, k) {
		for _, v1 := range v {
			print(v1, " ")
		}
		println()
	}
}
