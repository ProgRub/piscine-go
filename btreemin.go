package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	copia := root
	for copia.Left != nil {
		copia = copia.Left
	}
	return copia
}
