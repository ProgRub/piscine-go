package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	novo := &NodeI{Data: data_ref, Next: nil}
	if data_ref < l.Data {
		novo.Next = l
		l = novo
		return l
	}
	anterior := l
	itera := l.Next
	for itera != nil {
		if data_ref < itera.Data {
			novo.Next = itera
			anterior.Next = novo
			return l
		}
		anterior = anterior.Next
		itera = itera.Next
	}
	anterior.Next = novo
	return l
}
