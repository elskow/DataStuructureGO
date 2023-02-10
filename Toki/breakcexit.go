package main

import (
	"fmt"
)

func main() {
	var n int8
	fmt.Scan(&n)
	for i := 1; int8(i) <= n; i++ {
		if i%10== 0 {
			continue
		}else if i>41{
			fmt.Println("ERROR")
			break
		}else{
			fmt.Println(i)
		}
	}
}