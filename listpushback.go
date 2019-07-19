package piscine

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	novo := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = novo
		l.Tail = l.Head
		return
	}
	itera := l.Head
	for itera.Next != nil {
		itera = itera.Next
	}
	itera.Next = novo
	l.Tail = novo
}

func PrintList(l *List) {
	if l.Head != nil {
		itera := l.Head
		for itera != nil {
			fmt.Print(itera.Data, " -> ")
			itera = itera.Next
		}
		fmt.Println(nil)
	} else {
		fmt.Println("nil")
	}
}
