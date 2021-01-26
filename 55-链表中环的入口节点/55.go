func EntryNodeOfLoop(pHead *ListNode) *ListNode{
    m := make(map[*ListNode]int)
   
    //遍历链表，如果key不存在则赋值，已存在说明有环，遍历一遍后不存在则无环
    for pHead != nil {
        _, ok := m[pHead]
        if ok == false {
            m[pHead] = 1
            pHead = pHead.Next
        } else {
            return pHead
        }
    }
    
    return nil
}
