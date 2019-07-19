package piscine

func ListReverse(l *List) {
	if l.Head != nil {
		if ListSize(l) > 1 {
			dados := make([]interface{}, ListSize(l))
			itera := l.Head
			for itera != nil {
				dados = append(dados, itera.Data)
				itera = itera.Next
			}
			pos := 0
			final := &List{}
			final.Head = nil
			final.Tail = nil
			for pos < len(dados)/2 {
				ListPushFront(final, dados[pos])
				pos++
			}
			*l = *final
		}
	}
}
