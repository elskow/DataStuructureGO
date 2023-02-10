package main

import (
	"fmt"
)

func main() {
	var x,y int
	fmt.Scan(&x, &y)
	var matrices [101][101] int;
	
	for i:=0; i < int(x); i++{
		for j:=0; j < int(y); j++{
			fmt.Scan(&matrices[i][j])
		}
	}

	for i:=0; i < int(y); i++{
		for j:=int(x); j > 0; j--{
			fmt.Print(matrices[j-1][i], " ")
		}
		fmt.Println("")
	}
}