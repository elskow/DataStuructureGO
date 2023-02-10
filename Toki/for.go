package main

import (
	"fmt"
)

func main() {
	var T int16
	fmt.Scan(&T)

	result := 0

	for i:=0; i < int(T); i++{
		var temp int16
		fmt.Scan(&temp)
		result += int(temp)
	}
	fmt.Println(result)
}
