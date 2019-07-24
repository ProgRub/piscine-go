package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if root.Left == nil && root.Right == nil && root.Data == node.Data {
			root = nil
			return root
		}
		itera := root
		for itera != nil {
			if itera.Data == node.Data {
				if itera.Left != nil {
					aux := itera.Left
					for aux.Right != nil {
						aux = aux.Right
					}
					troca := aux.Data
					itera.Data = troca
					aux.Data = itera.Data
					aux = aux.Parent
					if aux.Left != nil && aux.Left.Data == node.Data {
						aux.Left = nil
					} else {
						aux.Right = nil
					}
				} else if itera.Right != nil {
					aux := itera.Right.Data
					itera.Data = aux
					tempLeft := itera.Right.Left
					tempRight := itera.Right.Right
					itera.Right = tempRight
					itera.Left = tempLeft
					if itera.Data == root.Data {
						root = itera
					}
				} else {
					p := itera.Parent
					if p.Left != nil && p.Left.Data == node.Data {
						p.Left = nil
					} else {
						p.Right = nil
					}
					return root
				}
				return root
			} else if node.Data < itera.Data {
				itera = itera.Left
			} else {
				itera = itera.Right
			}
		}
	}
	return root
}
