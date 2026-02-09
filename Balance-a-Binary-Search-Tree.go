var values []int

func buildBalancedBST(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNode{values[mid], nil, nil}
	root.Left = buildBalancedBST(start, mid-1)
	root.Right = buildBalancedBST(mid+1, end)

	return root
}

func getBSTValues(root *TreeNode) {
	if root.Left != nil {
		getBSTValues(root.Left)
	}
	values = append(values, root.Val)
	if root.Right != nil {
		getBSTValues(root.Right)
	}
}
func balanceBST(root *TreeNode) *TreeNode {
	values = make([]int, 0)
	// Inorder -> Left, Head, Right
	getBSTValues(root)
	fmt.Println(values)
	return buildBalancedBST(0, len(values)-1)

}