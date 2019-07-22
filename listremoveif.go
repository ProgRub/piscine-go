package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && CompStr(l.Head.Data, data_ref) {
		l.Head = l.Head.Next
	}
	if l.Head != nil {
		anterior := l.Head
		itera := anterior.Next
		for itera != nil {
			if CompStr(itera.Data, data_ref) {
				anterior.Next = itera.Next
				if itera.Next == nil {
					l.Tail = anterior
				}
			} else {
				anterior = anterior.Next
			}
			itera = itera.Next
		}
	} else {
		l.Tail = nil
	}
}
