package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if root == node {
			apagador := root
			if root.Right != nil {
				aux := root.Right
				if aux.Left == nil {
					root = root.Right
					root.Parent = nil
					apagador = nil
					return root
				}
				troca := BTreeMin(aux)
				auxiliar := root.Data
				root.Data = troca.Data
				troca.Data = auxiliar
				aux = troca.Parent
				aux.Left = troca.Left
				if aux.Left != nil {
					aux.Left.Parent = aux
				}
				apagador.Parent = nil
				apagador = nil
			} else if root.Left != nil {
				root = root.Left
				root.Parent = nil
				apagador.Parent = nil
				apagador = nil
			} else {
				root = nil
				apagador.Parent = nil
				apagador = nil
			}
			return root
		}
		itera := root
		apagador := itera
		for itera != nil {
			if itera == node && itera != root {
				apagador = itera
				if itera.Right != nil {
					aux := itera.Right
					if aux.Left == nil {
						aux := itera.Parent
						tempLeft := itera.Left
						if itera.Left != nil {
							itera.Left.Parent = itera.Left
						}
						itera = itera.Right
						esquerda := itera.Left
						for esquerda != nil && esquerda.Left != nil {
							esquerda = esquerda.Left
						}
						if esquerda == nil {
							esquerda = itera
						}
						esquerda.Left = tempLeft
						if itera != nil {
							itera.Parent = aux
						}
						if aux.Left != nil && aux.Left == node {
							aux.Left = itera
						} else {
							aux.Right = itera
						}
						apagador.Parent = nil
						apagador = nil
						return root
					}
					troca := BTreeMin(aux)
					auxiliar := itera.Data
					itera.Data = troca.Data
					troca.Data = auxiliar
					aux = troca.Parent
					aux.Left = troca.Right
					if aux.Left != nil {
						aux.Left.Parent = aux
					}
					troca.Parent = nil
					troca = nil
				} else if itera.Left != nil {
					aux := itera.Left.Data
					itera.Data = aux
					tempLeft := itera.Right.Left
					tempRight := itera.Right.Right
					if tempLeft != nil {
						tempLeft.Parent = itera
					}
					if tempRight != nil {
						tempRight.Parent = itera
					}
					itera.Right = tempRight
					itera.Left = tempLeft
					apagador.Parent = nil
					apagador = nil
				} else {
					p := itera.Parent
					if p.Left != nil && p.Left.Data == node.Data {
						p.Left = nil
					} else {
						p.Right = nil
					}
					apagador.Parent = nil
					apagador = nil
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
