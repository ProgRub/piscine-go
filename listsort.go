package piscine

import (
	"sort"
)

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	dados := make([]int, 0)
	itera := l
	for itera != nil {
		dados = append(dados, itera.Data)
		itera = itera.Next
	}
	sort.Ints(dados)
	l = nil
	for i := range dados {
		l = NodeIPushBack(l, dados[i])
	}
	return l
}

func NodeIPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data, Next: nil}
	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
