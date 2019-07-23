package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	copia := root
	for copia != nil {
		if elem < copia.Data {
			copia = copia.Left
		} else if elem > copia.Data {
			copia = copia.Right
		} else {
			return copia
		}
	}
	return copia
}
