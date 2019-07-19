package piscine

func ListReverse(l *List) {
	if l.Head != nil {
		itera := l.Head
		dados := make([]interface{}, ListSize(l))
		for itera != nil {
			dados = append(dados, itera.Data)
			itera = itera.Next
		}
		dados = Reverte(dados)
		pos := 0
		final := &List{}
		for pos < len(dados)/2 {
			ListPushBack(final, dados[pos])
			pos++
		}
		*l = *final
	}
}

func Reverte(d []interface{}) []interface{} {
	final := make([]interface{}, len(d))
	i := 0
	for index := len(d) - 1; index >= 0; index-- {
		final[i] = d[index]
		i++
	}
	return final
}
