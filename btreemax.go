package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	copia := root
	for copia.Right != nil {
		copia = copia.Right
	}
	return copia
}
