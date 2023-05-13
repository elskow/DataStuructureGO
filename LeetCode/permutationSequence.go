package main

func getPermutation(n int, k int) string {
	var res string
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	for i := 1; i < n; i++ {
		f := fact(n - i)
		idx := (k - 1) / f
		res += string('0' + nums[idx])
		nums = append(nums[:idx], nums[idx+1:]...)
		k -= idx * f
	}
	res += string('0' + nums[0])
	return res
}

func fact(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	n := 3
	k := 3
	println(getPermutation(n, k))
	n = 4
	k = 9
	println(getPermutation(n, k))
}
