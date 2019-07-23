package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	BTreeApplyByLevelAux(root, f, 0)
}

func BTreeApplyByLevelAux(root *TreeNode, f func(...interface{}) (int, error), level int) {
	h := BTreeLevelCount(root)
	for i := 1; i <= h; i++ {
		BTreeApplyToLevel(root, f, i)
	}
}

func BTreeApplyToLevel(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else {
		BTreeApplyToLevel(root.Left, f, level-1)
		BTreeApplyToLevel(root.Right, f, level-1)
	}
}
