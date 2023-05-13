package main

func insert(intervals [][]int, new []int) [][]int {
	var res [][]int
	var i int
	for i < len(intervals) && intervals[i][1] < new[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][0] <= new[1] {
		new[0] = min(new[0], intervals[i][0])
		new[1] = max(new[1], intervals[i][1])
		i++
	}
	res = append(res, new)
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	for _, interval := range insert(intervals, newInterval) {
		for _, num := range interval {
			print(num, " ")
		}
		println()
	}
}
