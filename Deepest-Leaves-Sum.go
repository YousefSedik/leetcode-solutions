func deepestLeavesSum(root *TreeNode) int {
	answer := 0
	// use BFS?
	endNode := *&TreeNode{Val: -1}
	q := make([]*TreeNode, 0, 60)
	q = append(q, root)
	q = append(q, &endNode)
	for len(q) != 1 && q[0] != &endNode {
		answer = 0
		for q[0].Val != -1 {
			answer += q[0].Val
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			q = q[1:]
		}
		q = q[1:]
		q = append(q, &endNode)
	}
	return answer
}
