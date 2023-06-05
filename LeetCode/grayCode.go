package main

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	result := []int{0, 1}
	for i := 1; i < n; i++ {
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, result[j]+(1<<uint(i)))
		}
	}

	return result
}

func main() {
	n := 2
	for _, num := range grayCode(n) {
		println(num)
	}
}
