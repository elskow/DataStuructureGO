package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var res []int
	res = append(res, 1)
	res = append(res, 2)
	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n-1]
}

func main() {
	n := 2
	println(climbStairs(n))
	n = 3
	println(climbStairs(n))
}
