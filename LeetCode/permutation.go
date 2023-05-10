package main

func permute(nums []int) [][]int {
	var result [][]int
	permuteHelper(nums, 0, &result)
	return result
}

func permuteHelper(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, nums...))
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permuteHelper(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func main() {
	nums := []int{1, 2, 3}
	for _, v := range permute(nums) {
		for _, vv := range v {
			print(vv, " ")
		}
		println()
	}
}
