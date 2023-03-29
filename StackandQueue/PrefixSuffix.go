//  konversi dari notasi bentuk infix menjadi bentuk postfix dan prefix menggunakan stack

package main

import (
	"fmt"
)

type Stack struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data string
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

func Predecessor(data string) int {
	if data == "+" || data == "-" {
		return 1
	} else if data == "*" || data == "/" {
		return 2
	} else if data == "^" {
		return 3
	} else {
		return 0
	}
}

func isoperator(data string) bool {
	if data == "+" || data == "-" || data == "*" || data == "/" || data == "^" {
		return true
	} else {
		return false
	}
}

func isoperand(data string) bool {
	if data != "+" && data != "-" && data != "*" && data != "/" && data != "^" && data != "(" && data != ")" {
		return true
	} else {
		return false
	}
}

func (list *Stack) Prefix(data string) {
	var node = Node{Data: data}
	if isoperand(data) {
		list.Push(&node)
	} else if isoperator(data) {
		list.Pop()
		list.Pop()
		list.Push(&node)
	}
}

func (list *Stack) Postfix(data string) {
	var node = Node{Data: data}
	if isoperand(data) {
		list.Push(&node)
	} else if isoperator(data) {
		list.Pop()
		list.Pop()
		list.Push(&node)
	}
}

func main() {
	infix := "A + B / C -  D *E"
	var list = Stack{}
	for i := 0; i < len(infix); i++ {
		if infix[i] != ' ' {
			var data = string(infix[i])
			if isoperand(data) {
				list.Postfix(data)
			} else if isoperator(data) {
				list.Postfix(data)
			}
		}
	}
	fmt.Print("Postfix : ")
	list.Read()
	fmt.Println()
	for i := 0; i < len(infix); i++ {
		if infix[i] != ' ' {
			var data = string(infix[i])
			if isoperand(data) {
				list.Prefix(data)
			} else if isoperator(data) {
				list.Prefix(data)
			}
		}
	}
	fmt.Print("Prefix : ")
	list.Read()
}