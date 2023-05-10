package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	result := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := mul + result[p2]
			result[p1] += sum / 10
			result[p2] = sum % 10
		}
	}
	str := ""
	for _, v := range result {
		if !(len(str) == 0 && v == 0) {
			str += strconv.Itoa(v)
		}
	}
	if len(str) == 0 {
		return "0"
	}
	return str
}

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
