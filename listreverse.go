package piscine

func ListReverse(l *List) {
	if l.Head != nil {
		if ListSize(l) > 2 {
			/* dados := make([]interface{}, 0)
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
			} */
			anterior := l.Head
			meio := l.Head.Next
			proximo := l.Head.Next.Next
			l.Head.Next = nil
			l.Head = l.Tail
			l.Tail = anterior
			for proximo != nil {
				meio.Next = anterior
				anterior = meio
				meio = proximo
				proximo = proximo.Next
			}
			meio.Next = anterior

		} else if ListSize(l) == 2 {
			l.Tail.Next = l.Head
			l.Head.Next = nil
		}
	}
}
