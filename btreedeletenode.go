package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if root == node {
			apagador := root
			if root.Left != nil {
				aux := root.Left
				if aux.Right == nil {
					root = root.Left
					root.Parent = nil
					apagador.Parent = nil
					apagador.Left = nil
					apagador.Right = nil
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
				apagador = nil
			} else if root.Right != nil {
				root = root.Right
				root.Parent = nil
				apagador = nil
			} else {
				root = nil
				apagador = nil
			}
			return root
		}
		itera := root
		apagador := itera
		apagador = apagador
		for itera != nil {
			if itera == node && itera != root {
				apagador = itera
				if itera.Left != nil {
					aux := itera.Left
					if aux.Right == nil {
						aux := itera.Parent
						tempRight := itera.Right
						if itera.Right != nil {
							itera.Right.Parent = itera.Left
						}
						itera = itera.Left
						direita := itera.Right
						for direita != nil && direita.Right != nil {
							direita = direita.Right
						}
						if direita == nil {
							direita = itera
						}
						direita.Right = tempRight
						if itera != nil {
							itera.Parent = aux
						}
						if aux.Left != nil && aux.Left == node {
							aux.Left = itera
						} else {
							aux.Right = itera
						}
						apagador = nil
						return root
					}
					troca := BTreeMax(aux)
					auxiliar := itera.Data
					itera.Data = troca.Data
					troca.Data = auxiliar
					aux = troca.Parent
					aux.Right = troca.Left
					if aux.Right != nil {
						aux.Right.Parent = aux
					}
					troca = nil
				} else if itera.Right != nil {
					aux := itera.Right.Data
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
					apagador = nil
				} else {
					p := itera.Parent
					if p.Left != nil && p.Left.Data == node.Data {
						p.Left = nil
					} else {
						p.Right = nil
					}
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
