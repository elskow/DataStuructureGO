package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membaca input hingga EOF
	scanner := bufio.NewScanner(os.Stdin)
	var arr []int
	for scanner.Scan() {
		var n int
		_, err := fmt.Sscan(scanner.Text(), &n)
		if err != nil {
			break
		}
		arr = append(arr, n)
	}

	// Membalik urutan array
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
