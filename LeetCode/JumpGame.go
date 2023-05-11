package main

func canJump(nums []int) bool {
	var max int
	for i, num := range nums {
		if i > max {
			return false
		}
		if i+num > max {
			max = i + num
		}
	}
	return true
}

func main() {
	println(canJump([]int{2, 3, 1, 1, 4}))
	println(canJump([]int{3, 2, 1, 0, 4}))
}
