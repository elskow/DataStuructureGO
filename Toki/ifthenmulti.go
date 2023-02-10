package main

import (
	"fmt"
)

func main() {
	var N int32
	fmt.Scan(&N)
	if N > 0 && N % 2 == 0{
		fmt.Println(N)
	} else{
		fmt.Println("")
	}
}