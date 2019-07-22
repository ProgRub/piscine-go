package piscine

func ListReverse(l *List) {
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
