package piscine

func ListPushFront(l *List, data interface{}) {
	novo := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = novo
		l.Tail = l.Head
		return
	}
	novo.Next = l.Head
	l.Head = novo
}
