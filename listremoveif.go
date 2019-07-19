package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head != nil {
		if CompStr(l.Head.Data, data_ref) {
			if l.Head.Next != nil {
				l.Head = l.Head.Next
				ListRemoveIf(l, data_ref)
			} else {
				l.Head = nil
				l.Tail = nil
				return
			}
		}
		itera := l.Head
		for itera.Next.Next != nil {
			if CompStr(itera.Next.Data, data_ref) {
				itera.Next = itera.Next.Next
			}
			itera = itera.Next
		}
		if CompStr(itera.Next.Data, data_ref) {
			itera.Next = nil
		}
	}
}
