package main

func pivotIndex(nums []int) int {
	var left, right int
	for i := 0; i < len(nums); i++ {
		right += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			left += nums[i-1]
		}
		right -= nums[i]
		if left == right {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	print(pivotIndex(nums))
}
