package main

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	dfs(nums, 0, path, &res)
	return res
}

func dfs(nums []int, index int, path []int, res *[][]int) {
	if index == len(nums) {
		var tmp []int
		tmp = append(tmp, path...)
		*res = append(*res, tmp)
		return
	}
	dfs(nums, index+1, path, res)
	path = append(path, nums[index])
	dfs(nums, index+1, path, res)
	path = path[:len(path)-1]
}

func main() {
	nums := []int{1, 2, 3}
	for _, v := range subsets(nums) {
		for _, v1 := range v {
			print(v1, " ")
		}
		println()
	}
	nums = []int{0}
	for _, v := range subsets(nums) {
		for _, v1 := range v {
			print(v1, " ")
		}
		println()
	}
}
