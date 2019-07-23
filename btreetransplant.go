package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	copia := root
	for copia != nil {
		if node.Data < copia.Data {
			copia = copia.Left
		} else if node.Data > copia.Data {
			copia = copia.Right
		} else {
			rplc.Parent = copia.Parent
			rplc.Left = copia.Left
			rplc.Right = copia.Right
			if copia != root {
				anterior := copia.Parent
				if anterior.Left.Data == copia.Data {
					anterior.Left = rplc
				} else {
					anterior.Right = rplc
				}
			} else {
				root = rplc
			}
			copia = nil
			return root
		}
	}
	return root
}
