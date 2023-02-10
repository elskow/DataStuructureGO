package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	for i := 1; i <= N; i++ {
		if i%K == 0 {
			fmt.Print("* ")
		} else {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("")
}