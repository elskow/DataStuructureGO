// Deskripsi
// N!, yaitu N faktorial, didefinisikan sebagai N × (N−1) × (N−2) × ... × 1.

// Pak Dengklek memberikan Anda sebuah bilangan bulat N (1 ≤ N ≤ 10.000). Hitunglah banyaknya 0 berurutan yang mengakhiri N!. Misalnya,

// 10! = 3.628.800, maka banyaknya 0 berurutan adalah 2.
// 8! = 40.320, maka banyaknya 0 berurutan adalah 1 (nol di tengah tidak dihitung).
// Format Masukan
// Baris pertama berisi sebuah bilangan bulat N.

// Format Keluaran
// Sebuah baris berisi sebuah bilangan bulat yaitu jawaban yang dimaksud.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var count int
	for i := 5; i <= n; i *= 5 {
		count += n / i
	}
	fmt.Println(count)
}