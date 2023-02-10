package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	} else if n % 2 == 0 {
		return isPowerOfTwo(n / 2)
	} else {
		return false
	}
}

func main() {
	var n uint32
	fmt.Scan(&n)
	if isPowerOfTwo(int(n)) {
		fmt.Println("ya")
	} else {
		fmt.Println("bukan")
	}
}
