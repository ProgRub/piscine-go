package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil && root == node {
		root = nil
		return root
	}
	copia := root
	var troca string
	for copia != nil {
		if copia.Data == node.Data {
			if copia.Left != nil {
				aux := copia.Left
				for aux.Right != nil {
					aux = aux.Right
				}
				troca = aux.Data
				copia.Data = troca
				aux.Data = copia.Data
				aux = aux.Parent
				if aux.Left.Data == node.Data {
					aux.Left = nil
				} else {
					aux.Right = nil
				}
			} else if copia.Right != nil {
				aux := copia.Right
				copia.Data = aux.Data
				copia.Right = nil
				if copia == root {
					root = copia
				}
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
