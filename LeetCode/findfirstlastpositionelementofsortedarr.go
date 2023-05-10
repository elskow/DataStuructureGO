package main

func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	result := -1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			result = mid
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

func findLast(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	result := -1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			result = mid
			start = mid + 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	for _, v := range searchRange(nums, target) {
		println(v)
	}
}
