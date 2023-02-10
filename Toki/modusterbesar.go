package main

import (
	"fmt"
)

func mode(arr [] uint32)uint32{
	frequency := make(map[uint32]uint32)
	for _, value := range arr {
		frequency[value]++
	}
	
	var mode uint32
	var maxFrequency uint32
	for key, value := range frequency {
		if value > maxFrequency {
			mode = key
			maxFrequency = value
		}
	}

	return mode
}

func main() {
	var t uint32
	arr := [] uint32{}
	fmt.Scan(&t)
	for i := 0; uint32(i) < t; i++ {
		var temp uint32
		fmt.Scan(&temp)
		arr = append(arr, temp)
	}
	fmt.Println(mode(arr))
}