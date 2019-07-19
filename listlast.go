package piscine

func ListLast(l *List) interface{} {
	if l.Head != nil {
		itera := l.Head
		for itera.Next != nil {
			itera = itera.Next
		}
		return itera.Data
	}
	return nil
}
