package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	itera := n1
	for itera.Next != nil {
		itera = itera.Next
	}
	itera.Next = n2
	return ListSort(n1)
}
