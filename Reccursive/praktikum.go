package main

import (
	"fmt"
	"time"
)

func pangkat(x uint, y uint) uint {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}

func optimizedPangkat(x uint, y uint) uint {
	if y == 0 {
		return 1
	} else if (y%2 == 0) {
		return optimizedPangkat(x*x, y/2)
	} else {
		return x * optimizedPangkat(x*x, (y-1)/2)
	}
}

func main() {
	fmt.Printf("Masukkan angka: ")
	var x uint
	fmt.Scan(&x)

	fmt.Printf("Masukkan pangkat: ")
	var y uint
	fmt.Scan(&y)

	start0 := time.Now()
	fmt.Printf("\n%d pangkat %d adalah %d (Menggunakan rekursif sederhana)\n", x, y, pangkat(x, y))
	elapsed := time.Since(start0)
	fmt.Printf("Waktu eksekusi: %s	\n\n", elapsed)

	start0 = time.Now()
	fmt.Printf("%d pangkat %d adalah %d (Menggunakan rekursif yang dioptimasi)\n", x, y, optimizedPangkat(x, y))
	elapsed = time.Since(start0)
	fmt.Printf("Waktu eksekusi: %s", elapsed)
}