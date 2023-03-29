// Create double linkedlist untuk menampung data mahasiswa

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Mahasiswa struct {
		Nama string
		NIM string
		IPK string
		Next *Mahasiswa
		Prev *Mahasiswa
}

type DoubleLinkedList struct {
		Head *Mahasiswa
		Tail *Mahasiswa
}

func (list *DoubleLinkedList) AddFirst(data *Mahasiswa) {
		if list.Head == nil {
				list.Head = data
				list.Tail = data
		} else {
				list.Head.Prev = data
				data.Next = list.Head
				list.Head = data
		}
}

func (list *DoubleLinkedList) AddLast(data *Mahasiswa) {
		if list.Head == nil {
				list.Head = data
				list.Tail = data
		} else {
				list.Tail.Next = data
				data.Prev = list.Tail
				list.Tail = data
		}
}

func (list *DoubleLinkedList) AddAfter(data *Mahasiswa, after *Mahasiswa) {
		if list.Head == nil {
				list.Head = data
				list.Tail = data
		} else {
				data.Next = after.Next
				data.Prev = after
				after.Next = data
		}
}

func (list *DoubleLinkedList) AddBefore(data *Mahasiswa, before *Mahasiswa) {
		if list.Head == nil {
				list.Head = data
				list.Tail = data
		} else {
				data.Next = before
				data.Prev = before.Prev
				before.Prev = data
		}
}

func (list *DoubleLinkedList) RemoveFirst() {
		if list.Head == nil {
				fmt.Println("List kosong")
		} else {
				list.Head = list.Head.Next
				list.Head.Prev = nil
		}
}

func (list *DoubleLinkedList) RemoveLast() {
		if list.Head == nil {
				fmt.Println("List kosong")
		} else {
				list.Tail = list.Tail.Prev
				list.Tail.Next = nil
		}
}

func (list *DoubleLinkedList) Remove(nama string) {
		if list.Head == nil {
				fmt.Println("List kosong")
		} else {
				temp := list.Head
				for temp.Nama != nama {
						temp = temp.Next
				}
				temp.Prev.Next = temp.Next
				temp.Next.Prev = temp.Prev
		}
}

func (list *DoubleLinkedList) Print() {
		if list.Head == nil {
				fmt.Println("List kosong")
		} else {
				temp := list.Head
				for temp != nil {
						fmt.Println(temp.Nama, temp.NIM, temp.IPK)
						temp = temp.Next
				}
		}
}

func (list *DoubleLinkedList) PrintReverse() {
		if list.Head == nil {
				fmt.Println("List kosong")
		} else {
				temp := list.Tail
				for temp != nil {
						fmt.Println(temp.Nama, temp.NIM, temp.IPK)
						temp = temp.Prev
				}
		}
}

func clearscreen() {
		fmt.Println("\033[H\033[2J")
}

func button() {
		fmt.Println("1. Tambah data mahasiswa dari depan")
		fmt.Println("2. Tambah data mahasiswa dari belakang")
		fmt.Println("3. Tambah data mahasiswa setelah data tertentu")
		fmt.Println("4. Tambah data mahasiswa sebelum data tertentu")
		fmt.Println("5. Hapus data mahasiswa dari depan")
		fmt.Println("6. Hapus data mahasiswa dari belakang")
		fmt.Println("7. Hapus data mahasiswa tertentu")
		fmt.Println("8. Cetak data mahasiswa dari depan")
		fmt.Println("9. Cetak data mahasiswa dari belakang")
		fmt.Println("10. Keluar")
}

func main() {
		var list DoubleLinkedList
		var pilih int
		var nama, nim, ipk string
		var data, after, before *Mahasiswa
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		for {
				clearscreen()
				button()
				fmt.Print("Pilih menu: ")
				fmt.Scanln(&pilih)
				switch pilih {
				case 1:
						fmt.Print("Masukkan nama: ")
						scanner.Scan()
						nama = scanner.Text()
						fmt.Print("Masukkan NIM: ")
						scanner.Scan()
						nim = scanner.Text()
						fmt.Print("Masukkan IPK: ")
						scanner.Scan()
						ipk = scanner.Text()
						data = &Mahasiswa{Nama: nama, NIM: nim, IPK: ipk}
						list.AddFirst(data)
				case 2:
						fmt.Print("Masukkan nama: ")
						scanner.Scan()
						nama = scanner.Text()
						fmt.Print("Masukkan NIM: ")
						scanner.Scan()
						nim = scanner.Text()
						fmt.Print("Masukkan IPK: ")
						scanner.Scan()
						ipk = scanner.Text()
						data = &Mahasiswa{Nama: nama, NIM: nim, IPK: ipk}
						list.AddLast(data)
				case 3:
						fmt.Print("Masukkan nama: ")
						scanner.Scan()
						nama = scanner.Text()
						fmt.Print("Masukkan NIM: ")
						scanner.Scan()
						nim = scanner.Text()
						fmt.Print("Masukkan IPK: ")
						scanner.Scan()
						ipk = scanner.Text()
						data = &Mahasiswa{Nama: nama, NIM: nim, IPK: ipk}
						fmt.Print("Masukkan nama setelah data yang akan ditambahkan: ")
						scanner.Scan()
						nama = scanner.Text()
						temp := list.Head
						for temp.Nama != nama {
								temp = temp.Next
						}
						after = temp
						list.AddAfter(data, after)
				case 4:
						fmt.Print("Masukkan nama: ")
						scanner.Scan()
						nama = scanner.Text()
						fmt.Print("Masukkan NIM: ")
						scanner.Scan()
						nim = scanner.Text()
						fmt.Print("Masukkan IPK: ")
						scanner.Scan()
						ipk = scanner.Text()
						data = &Mahasiswa{Nama: nama, NIM: nim, IPK: ipk}
						fmt.Print("Masukkan nama sebelum data yang akan ditambahkan: ")
						scanner.Scan()
						nama = scanner.Text()
						temp := list.Head
						for temp.Nama != nama {
								temp = temp.Next
						}
						before = temp
						list.AddBefore(data, before)
				case 5:
						list.RemoveFirst()
				case 6:
						list.RemoveLast()
				case 7:
						fmt.Print("Masukkan nama yang akan dihapus: ")
						scanner.Scan()
						nama = scanner.Text()
						list.Remove(nama)
				case 8:
						list.Print()
				case 9:
						list.PrintReverse()
				case 10:
						os.Exit(0)
				}
				fmt.Println("Tekan enter untuk melanjutkan")
				fmt.Scanln()
		}
}