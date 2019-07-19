package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l != nil {
		itera := l
		conta := 0
		for itera != nil && conta != pos {
			itera = itera.Next
			conta++
		}
		if itera != nil {
			return itera
		}
		return nil
	}
	return nil
}
