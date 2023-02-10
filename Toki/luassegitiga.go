package main

import (
	"fmt"
)

func main() {
	var A,T float32
	fmt.Scan(&A,&T)
	result := A*T/2
	fmt.Printf("%.2f\n",result)
}
