package main

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
}
