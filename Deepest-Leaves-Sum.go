func deepestLeavesSum(root *TreeNode) int {
	answer := 0
	// better if used BFS
	endNode := *&TreeNode{Val: -1}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	q = append(q, &endNode)
	for len(q) != 1 && q[0] != &endNode {
		// while q[0] != -1, pop q  and add children's to q
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
