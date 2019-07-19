package piscine

func ListLast(l *List) interface{} {
	if l.Head != nil {
		return l.Tail.Data
	}
	return nil
}
