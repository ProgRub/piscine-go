package piscine

func ListSize(l *List) int {
	conta := 0
	itera := l.Head
	for itera != nil {
		conta++
		itera = itera.Next
	}
	return conta
}
