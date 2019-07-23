package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil && root == node {
		root = nil
		return root
	}
	copia := root
	troca := &TreeNode{}
	for copia != nil {
		if copia.Data == node.Data {
			if copia.Left != nil {
				aux := copia.Left
				for aux.Right != nil {
					aux = aux.Right
				}
				troca.Data = aux.Data
				root = BTreeTransplant(root, copia, troca)
				copia = root
				for copia != nil {
					if copia.Data == troca.Data {
						copia.Left = nil
						break
					} else if troca.Data < copia.Data {
						copia = copia.Left
					} else {
						copia = copia.Right
					}
				}
			} else if copia.Right != nil {
				pai := copia.Parent
				pai.Right = copia.Right
			} else {
				p := copia.Parent
				if p.Left.Data == node.Data {
					p.Left = nil
				} else {
					p.Right = nil
				}
				return root
			}

			return root
		} else if node.Data < copia.Data {
			copia = copia.Left
		} else {
			copia = copia.Right
		}
	}
	return root
}
