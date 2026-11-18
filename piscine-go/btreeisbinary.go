package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBinary(root, "", "", false, false)
}

func isBinary(node *TreeNode, min string, max string, hasMin bool, hasMax bool) bool {
	if node == nil {
		return true
	}

	if hasMin && node.Data <= min {
		return false
	}
	if hasMax && node.Data >= max {
		return false
	}

	return isBinary(node.Left, min, node.Data, hasMin, true) &&
		isBinary(node.Right, node.Data, max, true, hasMax)
}
