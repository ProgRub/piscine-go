package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	itera := l.Head
	for itera != nil {
		if comp(itera.Data, ref) {
			return &itera.Data
		}
		itera = itera.Next
	}
	return nil
}
