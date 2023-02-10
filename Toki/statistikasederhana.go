package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	min := 0
	max := 0
	for i:=0; i < n; i++{
		var temp int
		fmt.Scan(&temp)
		if i == 0{
			min = temp
			max = temp
		}
		if temp < min{
			min = temp
		}
		if temp > max{
			max = temp
		}
	}
	fmt.Println(max, min)
}