package main

import (
	"fmt"
)

func main() {
	var A, B int32
	fmt.Scan(&A, &B)
	pembagian := A/B
	sisa := A%B
	fmt.Println("masing-masing ", pembagian)
	fmt.Println("bersisa ", sisa)
}
