package main

import (
	"fmt"
)

func main() {
	var n uint8
	fmt.Scan(&n)
	for i:=1; i <= int(n); i++{
		for j:=n; j >= 1; j--{
			if j > uint8(i){
				fmt.Print(" ")
			}else{
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}
}