package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	var (
		sign, res int
	)
	sign = 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}
	for dividend >= divisor {
		tmp := divisor
		i := 1
		for dividend >= tmp {
			dividend -= tmp
			res += i
			i <<= 1
			tmp <<= 1
		}
	}
	if sign < 0 {
		res = -res
	}
	if res > math.MaxInt32 {
		res = math.MaxInt32
	}
	return res
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
}
