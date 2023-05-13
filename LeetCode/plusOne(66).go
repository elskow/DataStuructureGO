package main

func plusOne(digits []int) []int {
	var res []int
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		res = append(res, sum%10)
		carry = sum / 10
	}
	if carry == 1 {
		res = append(res, 1)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func main() {
	digits := []int{1, 2, 3}
	println(plusOne(digits))
	digits = []int{4, 3, 2, 1}
	println(plusOne(digits))
}
