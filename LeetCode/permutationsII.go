package main

func permuteUnique(nums []int) [][]int {
	var result [][]int
	permuteUniqueHelper(nums, 0, &result)
	return result
}

func permuteUniqueHelper(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, nums...))
		return
	}
	for i := start; i < len(nums); i++ {
		if !containsDuplicate(nums, start, i) {
			nums[start], nums[i] = nums[i], nums[start]
			permuteUniqueHelper(nums, start+1, result)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
}

func containsDuplicate(nums []int, start, end int) bool {
	for i := start; i < end; i++ {
		if nums[i] == nums[end] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 1, 2}
	for _, v := range permuteUnique(nums) {
		for _, vv := range v {
			print(vv, " ")
		}
		println()
	}
}
