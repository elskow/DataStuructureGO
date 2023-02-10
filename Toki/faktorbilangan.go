package main

import (
	"fmt"
)

func factorization(n int){
	for i:=n ; i >= 1; i--{
		if n%i == 0{
			fmt.Println(i)
		}
	}
}


func main() {
	var n int
	fmt.Scan(&n)
	factorization(n)
}