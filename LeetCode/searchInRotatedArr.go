package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// find the pivot
	pivot := findPivot(nums)
	// fmt.Println("pivot: ", pivot)
	if pivot == -1 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}
	if nums[pivot] == target {
		return pivot
	}
	if nums[0] <= target {
		return binarySearch(nums, target, 0, pivot-1)
	}
	return binarySearch(nums, target, pivot+1, len(nums)-1)
}

func findPivot(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if mid < end && nums[mid] > nums[mid+1] {
			return mid
		}
		if mid > start && nums[mid] < nums[mid-1] {
			return mid - 1
		}
		if nums[mid] <= nums[start] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func binarySearch(nums []int, target int, start int, end int) int {
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
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	println(search(nums, target))
}
