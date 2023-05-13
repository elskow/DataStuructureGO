package main

func sortColors(nums []int) {
	var left, right int = 0, len(nums) - 1
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	for _, v := range nums {
		print(v, " ")
	}
	nums = []int{2, 0, 1}
	sortColors(nums)
	for _, v := range nums {
		print(v, " ")
	}
}
