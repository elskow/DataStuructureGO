package main

import (
	"bufio"
	"fmt"
	"os"
)

type DataKTP struct {
	Nama string
	Alamat string
	GolDarah string
	Next *DataKTP
	Prev *DataKTP
}

type DoubleLinkedList struct {
	Head *DataKTP
	Tail *DataKTP
}

func (list *DoubleLinkedList) InputDataDepan(data *DataKTP) {
	if list.Head == nil {
		list.Head = data
		list.Tail = data
	} else {
		list.Head.Prev = data
		data.Next = list.Head
		list.Head = data
	}
}

func (list *DoubleLinkedList) InputDataBelakang(data *DataKTP) {
	if list.Head == nil {
		list.Head = data
		list.Tail = data
	} else {
		list.Tail.Next = data
		data.Prev = list.Tail
		list.Tail = data
	}
}

func (list *DoubleLinkedList) InputDataPosisi(data *DataKTP, posisi *DataKTP) {
	if list.Head == nil {
		list.Head = data
		list.Tail = data
	} else {
		data.Next = posisi.Next
		data.Prev = posisi
		posisi.Next = data
	}
}

func (list *DoubleLinkedList) HapusDataDepan() {
	if list.Head == nil {
		fmt.Println("Data Kosong")
	} else {
		list.Head = list.Head.Next
		list.Head.Prev = nil
	}
}

func (list *DoubleLinkedList) HapusDataBelakang() {
	if list.Head == nil {
		fmt.Println("Data Kosong")
	} else {
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
	}
}

func (list *DoubleLinkedList) HapusDataPosisi(posisi *DataKTP) {
	if list.Head == nil {
		fmt.Println("Data Kosong")
	} else {
		posisi.Prev.Next = posisi.Next
		posisi.Next.Prev = posisi.Prev
	}
}

func (list *DoubleLinkedList) OutputData() {
	if list.Head == nil {
		fmt.Println("Data Kosong")
	} else {
		temp := list.Head
		for temp != nil {
			fmt.Println("Nama\t: ", temp.Nama)
			fmt.Println("Alamat\t: ", temp.Alamat)
			fmt.Println("GolDarah\t: ", temp.GolDarah)
			fmt.Println("====================================")
			temp = temp.Next
		}
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ConfirmationButton(){
	fmt.Println("Press Enter to Continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	var list DoubleLinkedList
	var input int
	var nama, alamat, golDarah string
	var data *DataKTP
	var posisi *DataKTP
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearScreen()
		fmt.Println("====================================")
		fmt.Println("Program Double Linked List")
		fmt.Println("====================================")
		fmt.Println("1. Input Data Depan")
		fmt.Println("2. Input Data Belakang")
		fmt.Println("3. Input Data Posisi")
		fmt.Println("4. Hapus Data Depan")
		fmt.Println("5. Hapus Data Belakang")
		fmt.Println("6. Hapus Data Posisi")
		fmt.Println("7. Output Data")
		fmt.Println("8. Exit")
		fmt.Println("====================================")
		fmt.Print("Pilih : ")
		fmt.Scanln(&input)
		switch input {
		case 1:
			ClearScreen()
			fmt.Println("====================================")
			fmt.Println("Input Data Depan")
			fmt.Println("====================================")
			fmt.Print("Nama\t: ")
			scanner.Scan()
			nama = scanner.Text()
			fmt.Print("Alamat\t: ")
			scanner.Scan()
			alamat = scanner.Text()
			fmt.Print("GolDarah\t: ")
			scanner.Scan()
			golDarah = scanner.Text()
			data = &DataKTP{nama, alamat, golDarah, nil, nil}
			list.InputDataDepan(data)
			ConfirmationButton()
		case 2:
			ClearScreen()
			fmt.Println("====================================")
			fmt.Println("Input Data Belakang")
			fmt.Println("====================================")
			fmt.Print("Nama\t: ")
			scanner.Scan()
			nama = scanner.Text()
			fmt.Print("Alamat\t: ")
			scanner.Scan()
			alamat = scanner.Text()
			fmt.Print("GolDarah\t: ")
			scanner.Scan()
			golDarah = scanner.Text()
			data = &DataKTP{nama, alamat, golDarah, nil, nil}
			list.InputDataBelakang(data)
			ConfirmationButton()
		case 3:
			ClearScreen()
			fmt.Println("====================================")
			fmt.Println("Input Data Posisi")
			fmt.Println("====================================")
			fmt.Print("Nama\t: ")
			scanner.Scan()
			nama = scanner.Text()
			fmt.Print("Alamat\t: ")
			scanner.Scan()
			alamat = scanner.Text()
			fmt.Print("GolDarah\t: ")
			scanner.Scan()
			golDarah = scanner.Text()
			data = &DataKTP{nama, alamat, golDarah, nil, nil}
			fmt.Print("Posisi\t: ")
			scanner.Scan()
			nama = scanner.Text()
			posisi = list.Head
			for posisi.Nama != nama {
				posisi = posisi.Next
			}
			list.InputDataPosisi(data, posisi)
			ConfirmationButton()
		case 4:
			ClearScreen()
			fmt.Println("====================================")
			fmt.Println("Hapus Data Depan")
			fmt.Println("====================================")
			list.HapusDataDepan()
			ConfirmationButton()
		case 5:
			ClearScreen()
			fmt.Println("====================================")
			fmt.Println("Hapus Data Belakang")
			fmt.Println("====================================")
			list.HapusDataBelakang()
			ConfirmationButton()
		case 6:
			ClearScreen()
			fmt.Println("====================================")
			fmt.Println("Hapus Data Posisi")
			fmt.Println("====================================")
			fmt.Print("Posisi\t: ")
			scanner.Scan()
			nama = scanner.Text()
			posisi = list.Head
			for posisi.Nama != nama {
				posisi = posisi.Next
			}
			list.HapusDataPosisi(posisi)
			ConfirmationButton()
		case 7:
			ClearScreen()
			fmt.Println("====================================")
			fmt.Println("Output Data")
			fmt.Println("====================================")
			list.OutputData()
			ConfirmationButton()
		case 8:
			os.Exit(0)
		default:
			fmt.Println("Pilihan Tidak Tersedia")
			ConfirmationButton()
		}
	}
}