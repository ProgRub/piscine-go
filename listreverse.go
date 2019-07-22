package piscine

func ListReverse(l *List) {
	if l.Head != nil {
		if ListSize(l) > 2 {
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
			aux := l.Head
			l.Head = l.Tail
			l.Tail = aux
		}
	}
}

/* func ListReverse(l *List) {
	anterior := &NodeL{}
	anterior = nil
	meio := l.Head
	proximo := &NodeL{}
	proximo = nil
	l.Head = l.Tail
	l.Tail = meio
	for meio != nil {
		proximo = meio.Next
		meio.Next = anterior
		anterior = meio
		meio = proximo
	}
}
*/
