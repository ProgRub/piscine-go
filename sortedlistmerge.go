package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	itera := n2
	for itera != nil {
		n1 = SortListInsert(n1, itera.Data)
		itera = itera.Next
	}
	return n1
}
