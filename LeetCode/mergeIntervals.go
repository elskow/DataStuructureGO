package main

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	var result [][]int
	var start, end int
	for i, interval := range intervals {
		if i == 0 {
			start, end = interval[0], interval[1]
			continue
		}
		if interval[0] > end {
			result = append(result, []int{start, end})
			start, end = interval[0], interval[1]
		} else if interval[1] > end {
			end = interval[1]
		}
	}
	result = append(result, []int{start, end})
	return result
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	println(merge(intervals))
}
