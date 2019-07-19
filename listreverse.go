package piscine

func ListReverse(l *List) {
	if l.Head != nil {
		if ListSize(l) > 1 {
			dados := make([]interface{}, 0)
			itera := l.Head
			for itera != nil {
				dados = append(dados, itera.Data)
				itera = itera.Next
			}
			pos := 0
			ListClear(l)
			for pos < len(dados) {
				ListPushFront(l, dados[pos])
				pos++
			}
		}
	}
}
