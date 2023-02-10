package main

import (
	"fmt"
)

func main() {
	var N int32
	fmt.Scan(&N)
	if N > 0{
		fmt.Println("positif")
	} else if N == 0{
		fmt.Println("nol")
	} else{
		fmt.Println("negatif")
	}
}