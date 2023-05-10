package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var jumps int = 0
	var maxReach int = 0
	var steps int = 0
	for i := 0; i < len(nums)-1; i++ {
		maxReach = max(maxReach, i+nums[i])
		if i == steps {
			jumps++
			steps = maxReach
		}
	}
	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	println(jump(nums))
}
