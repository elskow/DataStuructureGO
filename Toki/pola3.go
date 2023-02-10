package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	temp := 0
	for i := 1; i <= N; i++ {
		for j:=0; j<i;j++{
			if temp>9{
				temp = 0
			}
			fmt.Print(temp)
			temp++
		}
		fmt.Println("")
	}
}