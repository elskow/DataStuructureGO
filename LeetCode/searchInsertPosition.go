package main

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	println(searchInsert(nums, target))
}
