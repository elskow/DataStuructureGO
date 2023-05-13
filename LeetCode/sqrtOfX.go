package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var left, right, res int = 1, x, 0
	for left <= right {
		mid := (left + right) / 2
		if mid <= x/mid {
			left = mid + 1
			res = mid
		} else {
			right = mid - 1
		}
	}
	return res
}

func main() {
	x := 4
	println(mySqrt(x))
	x = 8
	println(mySqrt(x))
}
