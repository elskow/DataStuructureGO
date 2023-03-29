// Tugas Praktikum :

// Buat program Stack yang diimplementasikan menggunakan linkedlist untuk memasukkan,
// menghapus dan membaca data mahasiswa. field dari data mahasiswa bebas.

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

type Stack struct {
	Head *Mahasiswa
	Tail *Mahasiswa
}

func (list *Stack) Push(data *Mahasiswa) {
	if list.Head == nil {
		list.Head = data
		list.Tail = data
	} else {
		list.Head.Prev = data
		data.Next = list.Head
		list.Head = data
	}
}

func (list *Stack) Pop() {
	if list.Head == nil {
		fmt.Println("Data Kosong")
	} else {
		list.Head = list.Head.Next
		list.Head.Prev = nil
	}
}

func (list *Stack) Read() {
	if list.Head == nil {
		fmt.Println("Data Kosong")
	} else {
		var bantu = list.Head
		for bantu != nil {
			fmt.Println("Nama : ", bantu.Nama)
			fmt.Println("NIM : ", bantu.NIM)
			fmt.Println("IPK : ", bantu.IPK)
			bantu = bantu.Next
		}
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ConfirmationButton() {
	var input string
	fmt.Println("Press Enter to Continue")
	fmt.Scanln(&input)
}

func main() {
	var list Stack
	var input int
	scanner := bufio.NewScanner(os.Stdin)
	for input != 4 {
		ClearScreen()
		fmt.Println("====================================")
		fmt.Println("1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Read")
		fmt.Println("4. Exit")
		fmt.Println("====================================")
		fmt.Print("Input : ")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &input)
		switch input {
		case 1:
			var mahasiswa Mahasiswa
			fmt.Print("Nama : ")
			scanner.Scan()
			mahasiswa.Nama = scanner.Text()
			fmt.Print("NIM : ")
			scanner.Scan()
			mahasiswa.NIM = scanner.Text()
			fmt.Print("IPK : ")
			scanner.Scan()
			mahasiswa.IPK = scanner.Text()
			list.Push(&mahasiswa)
			ConfirmationButton()
		case 2:
			list.Pop()
			ConfirmationButton()
		case 3:
			list.Read()
			ConfirmationButton()
		}
	}
}