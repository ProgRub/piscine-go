package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	Left := BTreeLevelCount(root.Left)
	Right := BTreeLevelCount(root.Right)
	if Left > Right {
		return Left + 1
	}
	return Right + 1
}
