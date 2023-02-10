package main

import (
	"fmt"
	"math"
)

func main() {
	var X1, Y1, X2, Y2 int64
	fmt.Scan(&X1, &Y1, &X2, &Y2)
	fmt.Println(int64(math.Abs(float64(X1-X2)) + math.Abs(float64(Y1-Y2))))
}