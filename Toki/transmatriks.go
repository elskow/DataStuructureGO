package main

import (
	"fmt"
)

func main() {
	var(
		A int8
		B int8
		C int8
		D int8
		E int8
		F int8
		G int8
		H int8
		I int8
	)
	fmt.Scan(&A,&B,&C)
	fmt.Scan(&D,&E,&F)
	fmt.Scan(&G,&H,&I)

	fmt.Println(A,D,G)
	fmt.Println(B,E,H)
	fmt.Println(C,F,I)
}