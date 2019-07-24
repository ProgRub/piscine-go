package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if root.Data == node.Data {
			if root.Left != nil {
				aux := root.Left
				if aux.Right == nil {
					root = root.Left
					root.Parent = nil
					return root
				}
				troca := BTreeMax(aux)
				auxiliar := root.Data
				root.Data = troca.Data
				troca.Data = auxiliar
				aux = troca.Parent
				aux.Right = troca.Left
				if aux.Right != nil {
					aux.Right.Parent = aux
				}
			} else if root.Right != nil {
				root = root.Right
				root.Parent = nil
			} else {
				root = nil
			}
			return root
		}
		itera := root
		for itera != nil {
			if itera.Data == node.Data && itera != root {
				if itera.Left != nil {
					aux := itera.Left
					if aux.Right == nil {
						aux := itera.Parent
						itera = itera.Left
						if itera != nil {
							itera.Parent = aux
						}
						aux.Left = itera
						return root
					}
					troca := BTreeMax(aux)
					auxiliar := itera.Data
					itera.Data = troca.Data
					troca.Data = auxiliar
					aux = troca.Parent
					aux.Right = troca.Left
					aux.Right.Parent = aux
				} else if itera.Right != nil {
					aux := itera.Right.Data
					itera.Data = aux
					tempLeft := itera.Right.Left
					tempRight := itera.Right.Right
					tempLeft.Parent = itera
					tempRight.Parent = itera
					itera.Right = tempRight
					itera.Left = tempLeft
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
