func deleteMiddle(head *ListNode) *ListNode {
    counter := 0
    CountingNode := head;
    for CountingNode != nil {
		counter++
		CountingNode = CountingNode.Next
	}

	if counter == 1 {
		return nil
	}
	if counter == 2 {
		head.Next = nil
		return head
	}
	firstNode := head
	secondNode := head
	middleIndex := counter / 2 
	for i := 0; i < middleIndex; i++ {
		firstNode = secondNode
		secondNode = secondNode.Next
	
	}
	firstNode.Next = secondNode.Next
	return head
}
