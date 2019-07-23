package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		novo := &TreeNode{Data: data, Left: nil, Right: nil, Parent: nil}
		root = novo
	} else if data < root.Data {
		if root.Left == nil {
			novo := &TreeNode{Data: data, Left: nil, Right: nil, Parent: root}
			root.Left = novo
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			novo := &TreeNode{Data: data, Left: nil, Right: nil, Parent: root}
			root.Right = novo
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}
