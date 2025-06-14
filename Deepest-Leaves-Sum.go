/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func deepestLeavesSum(root *TreeNode) int {
	// get the max depth
	maxDepth := 1
	curDepth := 1
	var getMaxDepth func(node *TreeNode)
	getMaxDepth = func(node *TreeNode) {
		if node.Left != nil {
			curDepth++
			maxDepth = max(maxDepth, curDepth)
			getMaxDepth(node.Left)
			curDepth--
		}
		if node.Right != nil {
			curDepth++
			maxDepth = max(maxDepth, curDepth)
			getMaxDepth(node.Right)
			curDepth--
		}
	}
	getMaxDepth(root)
	// get the sum of the nodes where depth = max 
	answer := 0
	curDepth = 1
	var getSum func(node *TreeNode)
	getSum = func(node *TreeNode) {
		if curDepth == maxDepth{
			answer += node.Val
		}
		if node.Left != nil {
			curDepth++
			getSum(node.Left)
			curDepth--
		}
		if node.Right != nil {
			curDepth++
			getSum(node.Right)
			curDepth--
		}
	}
    
    getSum(root)
	return answer
}

