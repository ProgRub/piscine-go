package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	novo := &NodeI{Data: data_ref, Next: nil}
	if data_ref < l.Data {
		novo.Next = l
		l = novo
		return l
	}
	itera := l
	for itera.Next.Next != nil {
		if data_ref < itera.Next.Data {
			novo.Next = itera.Next
			itera.Next = novo
			return l
		}
		itera = itera.Next
	}
	if data_ref < itera.Next.Data {
		novo.Next = itera.Next
		itera.Next = novo
	}
	return l
}
