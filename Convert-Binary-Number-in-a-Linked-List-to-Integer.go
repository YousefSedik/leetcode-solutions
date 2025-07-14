/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	size := -1
	node := head
	for node != nil{
		size ++
		node = node.Next
	}
	res := 0.0
	node = head
	for node != nil{
		if node.Val == 1{
			res += math.Pow(2.0, float64(size))
		}
        node = node.Next
		size--
	}
	return int(res)
}
