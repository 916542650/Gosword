
// 解法三：hashmap
func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
    // write code here
	m := make(map[*ListNode]bool)
	for pHead1 != nil {
		m[pHead1] = true
		pHead1 = pHead1.Next
	}

	for pHead2 != nil {
		if _, ok := m[pHead2]; ok {
			return pHead2
		}
		pHead2 = pHead2.Next
	}
	return pHead2
}
