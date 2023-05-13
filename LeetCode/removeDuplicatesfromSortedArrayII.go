package main

func removeDuplicates(nums []int) int {
	var left, right int = 0, 0
	for right < len(nums) {
		if right < len(nums)-2 && nums[right] == nums[right+2] {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}
	}
	return left
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	println(removeDuplicates(nums))
}
