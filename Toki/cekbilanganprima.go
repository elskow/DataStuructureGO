package main

import (
	"fmt"
)

func isPrime(n uint32) bool{
	if n == 1 || n == 0{
		return false
	}
	for i:=uint32(2); i*i <= n; i++{
		if n%i == 0{
			return false
		}
	}
	return true	
}

func main() {
	var t uint32
	fmt.Scan(&t)
	for i:=uint32(0); i < t; i++{
		var n uint32
		fmt.Scan(&n)
		if isPrime(n){
			fmt.Println("YA")
		}else{
			fmt.Println("BUKAN")
		}
	}
}
