package main

func majorityElement(nums []int) int {
	var (
		majority int
		count    int
	)
	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if majority == num {
			count++
		} else {
			count--
		}
	}
	return majority
}

func main() {
	nums := []int{3, 2, 3}
	println(majorityElement(nums))
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	println(majorityElement(nums))
}
