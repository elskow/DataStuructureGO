// Decimal to Binary using Stack

package main

import (
	"fmt"
)

type Stack struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func (list *Stack) Push(data *Node) {
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
			fmt.Print(bantu.Data)
			bantu = bantu.Next
		}
	}
}

func main() {
	var list = Stack{}
	var input int
	fmt.Print("Masukkan Angka Desimal : ")
	fmt.Scan(&input)
	for input > 0 {
		var data = Node{Data: input % 2}
		list.Push(&data)
		input = input / 2
	}
	fmt.Print("Hasil Konversi : ")
	list.Read()
}
