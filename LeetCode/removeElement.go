package main

func removeElement(nums []int, val int) int {
	var (
		i, j int
	)
	for j < len(nums) {
		if nums[j] == val {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	println(removeElement(nums, val))
	for _, v := range nums {
		print(v, " ")
	}
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	println(removeElement(nums, val))
	for _, v := range nums {
		print(v, " ")
	}
}
